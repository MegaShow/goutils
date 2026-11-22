package caches

import (
	"context"
	"time"
)

// Option is option for cache.
//
// 缓存配置选项.
type Option[K comparable, V any] func(opts *Options[K, V])

// Options is the options for cache.
//
// 缓存配置选项.
type Options[K comparable, V any] struct {
	// CleanDuration is the duration of cache item to be cleaned asynchronously,
	// if not set, no background cleanup is performed.
	//
	// 缓存数据清理时间间隔, 如果未设定则不进行后台清理
	CleanDuration time.Duration

	// Expiration is the duration of cache item to be expired.
	//
	// 缓存数据过期时间.
	Expiration time.Duration

	// LoadFunc is the function to load cache item when item is not found or expired.
	//
	// 缓存数据加载函数, 当缓存数据不存在或已过期时调用.
	LoadFunc func(ctx context.Context, key K) (V, error)

	// RemoveFunc is called when an entry is removed.
	//
	// 条目移除回调, 实际替换、逐出或关闭清空时调用, 仅过期不触发
	RemoveFunc func(key K, value V)

	// Size is the size of cache, if not set, the cache will not be limited.
	//
	// 缓存容量, 如果未设定将不限制缓存容量.
	Size int
}

func applyOptions[K comparable, V any](o *Options[K, V], opts ...Option[K, V]) *Options[K, V] {
	if o == nil {
		o = &Options[K, V]{}
	}
	for _, fn := range opts {
		fn(o)
	}
	return o
}

// WithCleanDuration sets the duration of clean expired items asynchronously,
// if not set, no background cleanup is performed.
//
// 设置缓存数据清理时间间隔, 如果未设定则不进行后台清理
func WithCleanDuration[K comparable, V any](cleanDuration time.Duration) Option[K, V] {
	return func(opts *Options[K, V]) {
		if cleanDuration <= 0 {
			panic("clean duration cannot be negative")
		}
		opts.CleanDuration = cleanDuration
	}
}

// WithExpiration sets the expire duration of cache item, expiration cannot be negative.
//
// 设置缓存数据过期时间.
func WithExpiration[K comparable, V any](expiration time.Duration) Option[K, V] {
	return func(opts *Options[K, V]) {
		if expiration <= 0 {
			panic("expiration cannot be negative")
		}
		opts.Expiration = expiration
	}
}

// WithLoadFunc sets the loader of cache item when item is not found or expired.
//
// 设置缓存数据加载函数, 当缓存数据不存在或已过期时调用.
func WithLoadFunc[K comparable, V any](loadFunc func(ctx context.Context, key K) (V, error)) Option[K, V] {
	return func(opts *Options[K, V]) {
		opts.LoadFunc = loadFunc
	}
}

// WithRemoveFunc sets the callback for removed entries.
//
// 设置移除回调, 当缓存数据被移除时调用.
func WithRemoveFunc[K comparable, V any](removeFunc func(key K, value V)) Option[K, V] {
	return func(opts *Options[K, V]) {
		opts.RemoveFunc = removeFunc
	}
}

// WithSize sets the size of cache, if not set, the cache will not be limited.
//
// 设置缓存容量, 如果未设定将不限制缓存容量.
func WithSize[K comparable, V any](size int) Option[K, V] {
	return func(opts *Options[K, V]) {
		if size <= 0 {
			panic("size cannot be negative")
		}
		opts.Size = size
	}
}
