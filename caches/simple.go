package caches

import (
	"context"
	"sync"
	"time"

	"go.icytown.com/utils/usync"
)

type simpleItem[K comparable, V any] struct {
	key    K
	value  V
	expire time.Time
}

func (item *simpleItem[K, V]) isExpired(now *time.Time) bool {
	if item.expire.IsZero() {
		return false
	}
	if now != nil {
		return item.expire.Before(*now)
	}
	return item.expire.Before(time.Now())
}

type simpleCache[K comparable, V any] struct {
	opts  *Options[K, V]
	mu    sync.RWMutex
	sf    usync.Singleflight[K, V]
	items map[K]simpleItem[K, V]
	close chan struct{}
}

// NewSimpleCache creates a cache with simple implementation.
//
// 创建一个简单实现的缓存.
func NewSimpleCache[K comparable, V any](opts ...Option[K, V]) Cache[K, V] {
	c := &simpleCache[K, V]{
		opts:  applyOptions(nil, opts...),
		close: make(chan struct{}),
	}
	c.items = make(map[K]simpleItem[K, V], c.opts.Size)

	// 指定清理时间时, 创建 goroutine 定时清理缓存
	if c.opts.CleanDuration > 0 {
		go func() {
			maxSize := c.opts.Size
			if maxSize == 0 {
				maxSize = -1
			}
			ticker := time.NewTicker(c.opts.CleanDuration)
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:
					c.evict(maxSize)
				case <-c.close:
					return
				}
			}
		}()
	}

	return c
}

// Close stops cleaning and removes all entries.
// Callers must stop using the cache before Close; subsequent operations are undefined.
//
// 关闭缓存并移除所有条目; 调用前应停止使用缓存, 关闭后的操作行为未定义
func (c *simpleCache[K, V]) Close() error {
	c.mu.Lock()
	close(c.close)
	removed := make([]simpleItem[K, V], 0, len(c.items))
	for _, item := range c.items {
		removed = append(removed, item)
	}
	clear(c.items)
	c.mu.Unlock()
	c.notifyRemove(removed)
	return nil
}

// Get reads an item, or calls the configured loader if missing or expired.
//
// 读取数据, 缺失或过期时调用加载函数.
func (c *simpleCache[K, V]) Get(ctx context.Context, key K) (V, error) {
	c.mu.RLock()
	item, ok := c.items[key]
	c.mu.RUnlock()
	if ok && !item.isExpired(nil) {
		return item.value, nil
	}

	// 加载数据
	if c.opts.LoadFunc != nil {
		return c.sf.Do(key, func() (V, error) {
			value, err := c.opts.LoadFunc(ctx, key)
			if err == nil {
				_ = c.Set(ctx, key, value)
			}
			return value, err
		})
	}

	var zero V
	return zero, ErrNotFound
}

// Set writes item in cache by key.
//
// 根据 key 写入缓存数据.
func (c *simpleCache[K, V]) Set(_ context.Context, key K, value V) error {
	item := simpleItem[K, V]{key: key, value: value}
	if c.opts.Expiration > 0 {
		item.expire = time.Now().Add(c.opts.Expiration)
	}
	var removed []simpleItem[K, V]

	c.mu.Lock()
	old, ok := c.items[key]
	// 缓存已满时, 为新增条目腾出位置
	if !ok && c.opts.Size > 0 && len(c.items) >= c.opts.Size {
		removed = c.evictLocked(c.opts.Size - 1)
	}
	c.items[key] = item
	c.mu.Unlock()
	if ok {
		removed = append(removed, old)
	}
	c.notifyRemove(removed)
	return nil
}

func (c *simpleCache[K, V]) evict(maxSize int) {
	c.mu.Lock()
	removed := c.evictLocked(maxSize)
	c.mu.Unlock()
	c.notifyRemove(removed)
}

// evictLocked removes expired and excess entries and returns them for notification.
//
// 移除过期及超出容量的条目, 返回待通知的条目
func (c *simpleCache[K, V]) evictLocked(maxSize int) []simpleItem[K, V] {
	now := time.Now()
	var removed []simpleItem[K, V]

	// 清除过期数据
	for key, item := range c.items {
		if item.isExpired(&now) {
			delete(c.items, key)
			removed = append(removed, item)
		}
	}

	// 过期数据逐出还不够, 需要逐出未过期数据
	if maxSize >= 0 {
		for key, item := range c.items {
			if len(c.items) <= maxSize {
				break
			}
			delete(c.items, key)
			removed = append(removed, item)
		}
	}

	return removed
}

func (c *simpleCache[K, V]) notifyRemove(items []simpleItem[K, V]) {
	if c.opts.RemoveFunc != nil {
		for _, item := range items {
			c.opts.RemoveFunc(item.key, item.value)
		}
	}
}
