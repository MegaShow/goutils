// Package caches provides some cache implementation.
//
// 包 caches 提供了一些缓存的实现.
package caches

import (
	"context"
	"errors"
)

var (
	// ErrNotFound is the error returned when item not found in cache.
	//
	// 当缓存中找不到数据时返回的错误.
	ErrNotFound = errors.New("item not found")
)

// Cache is the universal interface for caches.
//
// 缓存通用接口.
type Cache[K comparable, V any] interface {
	Close() error
	Get(ctx context.Context, key K) (V, error)
	Set(ctx context.Context, key K, value V) error
}
