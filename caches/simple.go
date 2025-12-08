package caches

import (
	"context"
	"sync"
	"time"

	"go.icytown.com/utils/usync"
)

type simpleItem[V any] struct {
	value  V
	expire time.Time
}

func (item *simpleItem[V]) isExpired(now *time.Time) bool {
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
	items map[K]simpleItem[V]
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
	c.items = make(map[K]simpleItem[V], c.opts.Size)

	// 指定清理时间时, 创建 goroutine 定时清理缓存
	if c.opts.CleanDuration > 0 {
		go func() {
			ticker := time.NewTicker(c.opts.CleanDuration)
			for {
				select {
				case <-ticker.C:
					c.evict(c.opts.Size)
				case <-c.close:
					return
				}
			}
		}()
	}

	return c
}

// Close closes cache and stop clean goroutine.
//
// 关闭缓存和停止清理 goroutine.
func (c *simpleCache[K, V]) Close() error {
	close(c.close)
	return nil
}

// Get read item from cache by key, if sets loader func option and item is not found, it will call loader func and return.
//
// 根据 key 读取缓存数据, 如果设置了数据加载函数并且数据不存在, 则调用该函数并返回数据.
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
	item := simpleItem[V]{value: value}
	if c.opts.Expiration > 0 {
		item.expire = time.Now().Add(c.opts.Expiration)
	}

	// 未指定清理时间, 则在写入数据时检查是否需要清理
	if c.opts.CleanDuration == 0 {
		c.mu.RLock()
		cacheLen := len(c.items)
		c.mu.RUnlock()
		if c.opts.Size > 0 && cacheLen >= c.opts.Size {
			c.evict(c.opts.Size - 1)
		}
	}

	c.mu.Lock()
	c.items[key] = item
	c.mu.Unlock()
	return nil
}

func (c *simpleCache[K, V]) evict(maxSize int) {
	// 获取需要逐出的数据
	now := time.Now()
	evictKeys := make([]K, 0)
	c.mu.RLock()
	for key, item := range c.items {
		if item.isExpired(&now) {
			evictKeys = append(evictKeys, key)
		}
	}
	// 过期数据逐出还不够, 需要逐出未过期数据
	if maxSize > 0 && len(evictKeys) < len(c.items)-maxSize {
		for key, item := range c.items {
			if !item.isExpired(&now) {
				evictKeys = append(evictKeys, key)
				if len(evictKeys) >= len(c.items)-maxSize {
					break
				}
			}
		}
	}
	c.mu.RUnlock()
	if len(evictKeys) == 0 {
		return
	}

	// 清除数据
	c.mu.Lock()
	for _, key := range evictKeys {
		delete(c.items, key)
	}
	c.mu.Unlock()
}
