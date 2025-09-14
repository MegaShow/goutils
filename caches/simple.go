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
	baseCache[K, V]
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
		items: make(map[K]simpleItem[V]),
		close: make(chan struct{}),
	}
	c.initOpts(opts)

	// 创建 goroutine 定时清理缓存
	go func() {
		ticker := time.NewTicker(c.cleanDuration)
		for {
			select {
			case <-ticker.C:
				c.update()
			case <-c.close:
				return
			}
		}
	}()

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
	if c.loaderFunc != nil {
		value, err := c.sf.Do(key, func() (V, error) { return c.loaderFunc(ctx, key) })
		if err != nil {
			return value, err
		}
		_ = c.Set(ctx, key, value)
		return value, nil
	}

	var zero V
	return zero, ErrNotFound
}

// Set writes item in cache by key.
//
// 根据 key 写入缓存数据.
func (c *simpleCache[K, V]) Set(_ context.Context, key K, value V) error {
	item := simpleItem[V]{value: value}
	if c.expiration > 0 {
		item.expire = time.Now().Add(c.expiration)
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = item
	return nil
}

func (c *simpleCache[K, V]) update() {
	// 获取过期数据
	now := time.Now()
	expiredKeys := make([]K, 0)
	c.mu.RLock()
	for key, item := range c.items {
		if item.isExpired(&now) {
			expiredKeys = append(expiredKeys, key)
		}
	}
	c.mu.RUnlock()
	if len(expiredKeys) == 0 {
		return
	}

	// 清除过期数据
	c.mu.Lock()
	for _, key := range expiredKeys {
		item, ok := c.items[key]
		if ok && item.isExpired(&now) {
			delete(c.items, key)
		}
	}
	c.mu.Unlock()
}
