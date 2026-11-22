package caches

import (
	"context"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"go.icytown.com/utils/internal/assert"
)

func TestSimpleCache_Close(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	removed := make(map[string]int)
	cache := NewSimpleCache(
		WithCleanDuration[int, string](50*time.Millisecond),
		WithExpiration[int, string](10*time.Millisecond),
		WithRemoveFunc(func(_ int, value string) { removed[value]++ }),
	)
	assert.Nil(t, cache.Set(ctx, 1, "1"))
	assert.Nil(t, cache.Set(ctx, 2, "2"))

	// 测试关闭缓存后清空数据并通知所有移除条目
	assert.Nil(t, cache.Close())
	assert.Equal(t, 1, removed["1"])
	assert.Equal(t, 1, removed["2"])
	time.Sleep(100 * time.Millisecond)
	simple, ok := cache.(*simpleCache[int, string])
	assert.True(t, ok)
	assert.Equal(t, 0, len(simple.items))
	assert.Equal(t, 1, removed["1"])
	assert.Equal(t, 1, removed["2"])
}

func TestSimpleCache_Get(t *testing.T) {
	t.Parallel()
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
	cacheWithLoader := NewSimpleCache(WithLoadFunc(func(ctx context.Context, key int) (string, error) {
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

	// 测试访问过期条目不会移除数据或触发回调
	var removed atomic.Int32
	expiredCache, ok := NewSimpleCache(WithRemoveFunc(func(int, string) { removed.Add(1) })).(*simpleCache[int, string])
	assert.True(t, ok)
	defer func() { _ = expiredCache.Close() }()
	expiredCache.items[1] = simpleItem[int, string]{key: 1, value: "1", expire: time.Now().Add(-time.Second)}
	value, err = expiredCache.Get(ctx, 1)
	assert.Equal(t, ErrNotFound, err)
	assert.Zero(t, value)
	assert.Equal(t, 1, len(expiredCache.items))
	assert.Equal(t, int32(0), removed.Load())

	// 测试并发读取不需要写锁
	readerCache, ok := NewSimpleCache[int, string]().(*simpleCache[int, string])
	assert.True(t, ok)
	defer func() { _ = readerCache.Close() }()
	assert.Nil(t, readerCache.Set(ctx, 1, "1"))
	readerDone := make(chan struct{})
	readerCache.mu.RLock()
	go func() {
		defer close(readerDone)
		value, err := readerCache.Get(ctx, 1)
		assert.Nil(t, err)
		assert.Equal(t, "1", value)
	}()
	select {
	case <-readerDone:
	case <-time.After(3 * time.Second):
		t.Error("cache hit blocked by another reader")
	}
	readerCache.mu.RUnlock()
	<-readerDone

	// 测试加载替换旧值后, 移除回调可以读取新值
	started, resume := make(chan struct{}), make(chan struct{})
	loadDone := make(chan error, 1)
	var reentryCache Cache[int, string]
	reentryCache = NewSimpleCache(WithLoadFunc(func(context.Context, int) (string, error) {
		close(started)
		<-resume
		return "loaded", nil
	}), WithRemoveFunc(func(key int, value string) {
		if value == "old" {
			current, err := reentryCache.Get(ctx, key)
			assert.Nil(t, err)
			assert.Equal(t, "loaded", current)
		}
	}))
	defer func() { _ = reentryCache.Close() }()
	go func() { _, err := reentryCache.Get(ctx, 1); loadDone <- err }()
	<-started
	assert.Nil(t, reentryCache.Set(ctx, 1, "old"))
	close(resume)
	select {
	case err := <-loadDone:
		assert.Nil(t, err)
	case <-time.After(3 * time.Second):
		t.Fatal("loader removal callback deadlocked")
	}
}

func TestSimpleCache_Set(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	cache := NewSimpleCache[int, string]()
	defer func() { _ = cache.Close() }()
	assert.Nil(t, cache.Set(ctx, 1, "1"))
	value, err := cache.Get(ctx, 1)
	assert.Nil(t, err)
	assert.Equal(t, "1", value)

	// 测试替换旧值时通知移除, 回调在锁外执行
	removed := make(map[string]int)
	var replacementCache Cache[int, string]
	replacementCache = NewSimpleCache(WithRemoveFunc(func(key int, value string) {
		removed[value]++
		if value == "1" {
			current, err := replacementCache.Get(ctx, key)
			assert.Nil(t, err)
			assert.Equal(t, "2", current)
			assert.Nil(t, replacementCache.Set(ctx, 2, "3"))
		}
	}))
	defer func() { _ = replacementCache.Close() }()
	assert.Nil(t, replacementCache.Set(ctx, 1, "1"))
	setDone := make(chan struct{})
	go func() { _ = replacementCache.Set(ctx, 1, "2"); close(setDone) }()
	select {
	case <-setDone:
	case <-time.After(3 * time.Second):
		t.Fatal("replacement callback deadlocked")
	}
	assert.Equal(t, 1, removed["1"])
	assert.Equal(t, 0, removed["2"])

	// 测试容量为一时写入新值会逐出旧值
	capacityRemoved := make(map[string]int)
	capacityCache := NewSimpleCache(WithSize[int, string](1), WithRemoveFunc(func(_ int, value string) { capacityRemoved[value]++ }))
	defer func() { _ = capacityCache.Close() }()
	assert.Nil(t, capacityCache.Set(ctx, 1, "1"))
	assert.Nil(t, capacityCache.Set(ctx, 2, "2"))
	assert.Equal(t, 1, capacityRemoved["1"])
	assert.Equal(t, 0, capacityRemoved["2"])
	value, err = capacityCache.Get(ctx, 2)
	assert.Nil(t, err)
	assert.Equal(t, "2", value)

	// 测试配置后台清理后, 写入仍立即遵守容量上限
	limitedCache, ok := NewSimpleCache(WithSize[int, int](2), WithCleanDuration[int, int](time.Hour)).(*simpleCache[int, int])
	assert.True(t, ok)
	defer func() { _ = limitedCache.Close() }()
	assert.Nil(t, limitedCache.Set(ctx, 1, 1))
	assert.Nil(t, limitedCache.Set(ctx, 2, 2))
	assert.Nil(t, limitedCache.Set(ctx, 1, 3))
	valueAtTwo, err := limitedCache.Get(ctx, 2)
	assert.Nil(t, err)
	assert.Equal(t, 2, valueAtTwo)
	assert.Nil(t, limitedCache.Set(ctx, 3, 3))
	assert.Equal(t, 2, len(limitedCache.items))

	// 测试并发写入不会超过容量, 每个移除值仅回调一次
	var removedCount atomic.Int32
	concurrentCache, ok := NewSimpleCache(WithSize[int, int](2), WithRemoveFunc(func(int, int) { removedCount.Add(1) })).(*simpleCache[int, int])
	assert.True(t, ok)
	var writers sync.WaitGroup
	start := make(chan struct{})
	for key := range 100 {
		writers.Add(1)
		go func() {
			defer writers.Done()
			<-start
			_ = concurrentCache.Set(ctx, key, key)
		}()
	}
	close(start)
	writers.Wait()
	assert.Equal(t, 2, len(concurrentCache.items))
	assert.Equal(t, int32(98), removedCount.Load())
	assert.Nil(t, concurrentCache.Close())
	assert.Equal(t, int32(100), removedCount.Load())
}

func Test_simpleCache_evict(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	var removed atomic.Int32
	cache, ok := NewSimpleCache(WithRemoveFunc(func(int, string) { removed.Add(1) })).(*simpleCache[int, string])
	assert.True(t, ok)
	defer func() { _ = cache.Close() }()
	assert.Nil(t, cache.Set(ctx, 1, "1"))

	// 测试无限容量清理保留未过期条目
	cache.evict(-1)
	value, err := cache.Get(ctx, 1)
	assert.Nil(t, err)
	assert.Equal(t, "1", value)
	assert.Equal(t, int32(0), removed.Load())

	// 测试并发清理同一过期条目只回调一次
	item := cache.items[1]
	item.expire = time.Now().Add(-time.Second)
	cache.items[1] = item
	var wg sync.WaitGroup
	for range 30 {
		wg.Add(1)
		go func() { defer wg.Done(); cache.evict(-1) }()
	}
	wg.Wait()
	assert.Equal(t, int32(1), removed.Load())
	assert.Equal(t, 0, len(cache.items))

	// 测试后台定时清理触发移除回调
	backgroundRemoved := make(chan string, 1)
	backgroundCache := NewSimpleCache(WithCleanDuration[int, string](time.Millisecond), WithExpiration[int, string](time.Millisecond), WithRemoveFunc(func(_ int, value string) { backgroundRemoved <- value }))
	defer func() { _ = backgroundCache.Close() }()
	assert.Nil(t, backgroundCache.Set(ctx, 1, "1"))
	select {
	case value := <-backgroundRemoved:
		assert.Equal(t, "1", value)
	case <-time.After(3 * time.Second):
		t.Fatal("expired entry was not removed")
	}
}
