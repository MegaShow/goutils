package usync

import (
	"errors"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"go.icytown.com/utils/internal/assert"
)

func TestSingleflight_Do(t *testing.T) {
	t.Parallel()
	fnBuilder := func(key int, counter *int32) func() (string, error) {
		return func() (string, error) {
			atomic.AddInt32(counter, 1)
			if key < 0 {
				panic("negative key")
			} else if key == 0 {
				return "", errors.New("key is zero")
			}
			return strconv.Itoa(key), nil
		}
	}

	// 测试不同 key 的请求独立执行
	sf, counter, wg := Singleflight[int, string]{}, int32(0), sync.WaitGroup{}
	wg.Add(3)
	for i := range 3 {
		go func() {
			defer wg.Done()
			value, err := sf.Do(i+1, fnBuilder(i+1, &counter))
			assert.Equal(t, strconv.Itoa(i+1), value)
			assert.Nil(t, err)
		}()
	}
	wg.Wait()
	assert.Equal(t, int32(3), counter)

	// 测试请求结束后相同 key 可以再次加载
	sf, counter = Singleflight[int, string]{}, 0
	for range 3 {
		value, err := sf.Do(1, fnBuilder(1, &counter))
		assert.Equal(t, "1", value)
		assert.Nil(t, err)
	}
	assert.Equal(t, int32(3), counter)

	// 测试错误和 panic 不会阻止后续重试
	_, err := sf.Do(1, fnBuilder(0, &counter))
	assert.Equal(t, errors.New("key is zero"), err)
	_, err = sf.Do(1, fnBuilder(-1, &counter))
	assert.NotNil(t, err)
	assert.True(t, strings.HasPrefix(err.Error(), "panic: negative key\n\ngoroutine "))
	value, err := sf.Do(1, fnBuilder(1, &counter))
	assert.Nil(t, err)
	assert.Equal(t, "1", value)

	// 预置进行中的加载记录, 验证等待者复用结果且不执行自己的加载函数
	// 记录保留到等待者全部返回, 避免依赖 goroutine 的调度先后
	for _, failed := range []bool{false, true} {
		pending := &call[string]{}
		pending.wg.Add(1)
		cache := Singleflight[int, string]{m: map[int]*call[string]{1: pending}}
		var calls atomic.Int32
		ready, done := make(chan struct{}, 3), make(chan struct{}, 3)
		for range 3 {
			go func() {
				ready <- struct{}{}
				value, err := cache.Do(1, func() (string, error) {
					calls.Add(1)
					return "unexpected", nil
				})
				if failed {
					assert.Equal(t, "", value)
					assert.Equal(t, errors.New("load failed"), err)
				} else {
					assert.Equal(t, "shared", value)
					assert.Nil(t, err)
				}
				done <- struct{}{}
			}()
		}
		for range 3 {
			<-ready
		}
		if failed {
			pending.err = errors.New("load failed")
		} else {
			pending.value = "shared"
		}
		pending.wg.Done()
		for range 3 {
			<-done
		}
		assert.Equal(t, int32(0), calls.Load())
	}
}
