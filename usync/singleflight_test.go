package usync

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSingleflight_Do(t *testing.T) {
	fnBuilder := func(key int, counter *int32) func() (string, error) {
		return func() (string, error) {
			atomic.AddInt32(counter, 1)
			time.Sleep(50 * time.Millisecond) // 模拟耗时动作
			if key < 0 {
				panic("negative key")
			} else if key == 0 {
				return "", errors.New("key is zero")
			}
			return strconv.FormatInt(int64(key), 10), nil
		}
	}

	// 测试同一个 key 只发起一次请求
	sf, counter, wg := Singleflight[int, string]{}, int32(0), sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			value, err := sf.Do(1, fnBuilder(1, &counter))
			assert.Equal(t, "1", value)
			assert.Nil(t, err)
		}()
	}
	wg.Wait()
	assert.Equal(t, int32(1), counter)

	// 测试不同 key 发起多次请求
	sf, counter, wg = Singleflight[int, string]{}, int32(0), sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(idx int) {
			defer wg.Done()
			value, err := sf.Do(idx+1, fnBuilder(idx+1, &counter))
			assert.Equal(t, fmt.Sprint(idx+1), value)
			assert.Nil(t, err)
		}(i)
	}
	wg.Wait()
	assert.Equal(t, int32(3), counter)

	// 测试请求结束后相同 key 再次发起请求
	sf, counter, wg = Singleflight[int, string]{}, int32(0), sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		value, err := sf.Do(1, fnBuilder(1, &counter))
		assert.Equal(t, "1", value)
		assert.Nil(t, err)
	}
	assert.Equal(t, int32(3), counter)

	// 测试请求返回错误
	sf, counter, wg = Singleflight[int, string]{}, int32(0), sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			value, err := sf.Do(1, fnBuilder(0, &counter))
			assert.Equal(t, "", value)
			assert.Equal(t, errors.New("key is zero"), err)
		}()
	}
	wg.Wait()
	assert.Equal(t, int32(1), counter)

	// 测试请求 panic
	sf, counter, wg = Singleflight[int, string]{}, int32(0), sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			value, err := sf.Do(1, fnBuilder(-1, &counter))
			assert.Equal(t, "", value)
			assert.True(t, strings.HasPrefix(err.Error(), "panic: negative key\n\ngoroutine "))
		}()
	}
	wg.Wait()
	assert.Equal(t, int32(1), counter)
}
