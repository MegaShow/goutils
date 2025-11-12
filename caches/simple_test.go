package caches

import (
	"context"
	"strconv"
	"testing"
	"time"

	"go.icytown.com/utils/internal/assert"
)

func TestSimpleCache_Close(t *testing.T) {
	// 测试关闭缓存后不再清理过期数据
	ctx := context.Background()
	cache := NewSimpleCache(WithCleanDuration[int, string](50*time.Millisecond), WithExpiration[int, string](10*time.Millisecond))
	assert.Nil(t, cache.Set(ctx, 1, "1"))
	assert.Nil(t, cache.Close())
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, 1, len(cache.(*simpleCache[int, string]).items))
}

func TestSimpleCache_Get(t *testing.T) {
	ctx := context.Background()
	cache := NewSimpleCache(WithCleanDuration[int, string](50*time.Millisecond), WithExpiration[int, string](10*time.Millisecond))
	defer func() { _ = cache.Close() }()
	assert.Nil(t, cache.Set(ctx, 1, "1"))

	// 测试获取正常数据
	value, err := cache.Get(ctx, 1)
	assert.Nil(t, err)
	assert.Equal(t, "1", value)

	// 测试找不到正常数据
	value, err = cache.Get(ctx, 2)
	assert.Equal(t, ErrNotFound, err)
	assert.Zero(t, value)

	// 测试找不到过期数据
	time.Sleep(10 * time.Millisecond)
	value, err = cache.Get(ctx, 1)
	assert.Equal(t, ErrNotFound, err)
	assert.Zero(t, value)

	// 测试找不到已清除数据
	time.Sleep(50 * time.Millisecond)
	value, err = cache.Get(ctx, 1)
	assert.Equal(t, ErrNotFound, err)
	assert.Zero(t, value)

	// 测试数据加载
	cacheWithLoader := NewSimpleCache(WithLoaderFunc(func(ctx context.Context, key int) (string, error) {
		if key <= 0 {
			return "", ErrNotFound
		}
		return strconv.FormatInt(int64(key), 10), nil
	}))
	defer func() { _ = cacheWithLoader.Close() }()
	value, err = cacheWithLoader.Get(ctx, 0)
	assert.Equal(t, ErrNotFound, err)
	assert.Zero(t, value)
	value, err = cacheWithLoader.Get(ctx, 1)
	assert.Nil(t, err)
	assert.Equal(t, "1", value)
}

func TestSimpleCache_Set(t *testing.T) {
	ctx := context.Background()
	cache := NewSimpleCache[int, string]()
	defer func() { _ = cache.Close() }()
	assert.Nil(t, cache.Set(ctx, 1, "1"))
	value, err := cache.Get(ctx, 1)
	assert.Nil(t, err)
	assert.Equal(t, "1", value)
}
