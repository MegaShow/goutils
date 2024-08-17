// Package caches provides some cache implementation.
//
// 包 caches 提供了一些缓存的实现.
package caches

import (
	"context"
	"errors"
	"time"
)

var (
	// NotFound is the error returned when item not found in cache.
	//
	// 当缓存中找不到数据时返回的错误.
	NotFound = errors.New("item not found")
)

// Cache is the universal interface for caches.
//
// 缓存通用接口.
type Cache[K comparable, V any] interface {
	Close() error
	Get(ctx context.Context, key K) (V, error)
	Set(ctx context.Context, key K, value V) error
}

// Option is option for cache.
//
// 缓存配置选项.
type Option[K comparable, V any] func(c *baseCache[K, V])

// WithCleanDuration sets the duration of clean expired items, default by 1 minutes (maybe modify in the future)
//
// 设置过期数据清理的周期, 默认1分钟 (未来可能修改)
func WithCleanDuration[K comparable, V any](cleanDuration time.Duration) Option[K, V] {
	return func(c *baseCache[K, V]) {
		if cleanDuration <= 0 {
			panic("clean duration cannot be negative")
		}
		c.cleanDuration = cleanDuration
	}
}

// WithExpiration sets the expire duration of cache item, expiration cannot be negative.
//
// 设置缓存数据过期时间.
func WithExpiration[K comparable, V any](expiration time.Duration) Option[K, V] {
	return func(c *baseCache[K, V]) {
		if expiration <= 0 {
			panic("expiration cannot be negative")
		}
		c.expiration = expiration
	}
}

// WithLoaderFunc sets the loader of cache item when item is not found.
//
// 设置缓存数据加载函数.
func WithLoaderFunc[K comparable, V any](loaderFunc func(ctx context.Context, key K) (V, error)) Option[K, V] {
	return func(c *baseCache[K, V]) {
		c.loaderFunc = loaderFunc
	}
}
