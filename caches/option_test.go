package caches

import (
	"context"
	"errors"
	"testing"
	"time"

	"go.icytown.com/utils/internal/assert"
)

func TestWithCleanDuration(t *testing.T) {
	t.Parallel()
	opts := applyOptions(nil, WithCleanDuration[int, int](time.Minute))
	assert.Equal(t, time.Minute, opts.CleanDuration)
	opts = applyOptions(opts, WithCleanDuration[int, int](time.Hour))
	assert.Equal(t, time.Hour, opts.CleanDuration)

	// 测试零值和负值无效
	for _, value := range []time.Duration{0, -time.Second} {
		assert.Panics(t, func() { applyOptions(nil, WithCleanDuration[int, int](value)) })
	}
}

func TestWithExpiration(t *testing.T) {
	t.Parallel()
	opts := applyOptions(nil, WithExpiration[int, int](time.Minute))
	assert.Equal(t, time.Minute, opts.Expiration)
	opts = applyOptions(opts, WithExpiration[int, int](time.Hour))
	assert.Equal(t, time.Hour, opts.Expiration)

	// 测试零值和负值无效
	for _, value := range []time.Duration{0, -time.Second} {
		assert.Panics(t, func() { applyOptions(nil, WithExpiration[int, int](value)) })
	}
}

func TestWithLoadFunc(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	loadErr := errors.New("load failed")
	opts := applyOptions(nil, WithLoadFunc(func(received context.Context, key int) (int, error) {
		assert.Equal(t, ctx, received)
		if key < 0 {
			return 0, loadErr
		}
		return key + 1, nil
	}))
	value, err := opts.LoadFunc(ctx, 42)
	assert.Equal(t, 43, value)
	assert.Nil(t, err)
	value, err = opts.LoadFunc(ctx, -1)
	assert.Zero(t, value)
	assert.Equal(t, loadErr, err)

	// 测试清除已配置的加载函数
	opts = applyOptions(opts, WithLoadFunc[int, int](nil))
	assert.Nil(t, opts.LoadFunc)
}

func TestWithRemoveFunc(t *testing.T) {
	t.Parallel()
	calls := 0
	opts := applyOptions(nil, WithRemoveFunc(func(key int, value string) {
		calls++
		assert.Equal(t, 42, key)
		assert.Equal(t, "value", value)
	}))
	opts.RemoveFunc(42, "value")
	assert.Equal(t, 1, calls)

	// 测试清除已配置的移除回调
	opts = applyOptions(opts, WithRemoveFunc[int, string](nil))
	assert.Nil(t, opts.RemoveFunc)
}

func TestWithSize(t *testing.T) {
	t.Parallel()
	opts := applyOptions(nil, WithSize[int, int](1))
	assert.Equal(t, 1, opts.Size)
	opts = applyOptions(opts, WithSize[int, int](10))
	assert.Equal(t, 10, opts.Size)

	// 测试零值和负值无效
	for _, value := range []int{0, -1} {
		assert.Panics(t, func() { applyOptions(nil, WithSize[int, int](value)) })
	}
}

func Test_applyOptions(t *testing.T) {
	t.Parallel()
	opts := applyOptions[int, int](nil)
	assert.NotNil(t, opts)
	assert.Zero(t, *opts)

	// 测试配置组合、覆盖顺序及保留原配置
	opts = applyOptions(opts, WithSize[int, int](1), WithSize[int, int](2), WithExpiration[int, int](time.Minute))
	assert.Equal(t, 2, opts.Size)
	assert.Equal(t, time.Minute, opts.Expiration)
	updated := applyOptions(opts, WithCleanDuration[int, int](time.Hour))
	assert.Equal(t, opts, updated)
	assert.Equal(t, 2, updated.Size)
	assert.Equal(t, time.Minute, updated.Expiration)
	assert.Equal(t, time.Hour, updated.CleanDuration)
}
