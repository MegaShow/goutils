package usync

import (
	"fmt"
	"runtime/debug"
	"sync"
)

// Singleflight provides methods which can execute function once with same key and same time.
//
// 提供了相同 key 同一时间只执行一次函数的方法.
type Singleflight[K comparable, V any] struct {
	mu sync.Mutex
	m  map[K]*call[V]
}

type call[V any] struct {
	wg    sync.WaitGroup
	value V
	err   error
}

// Do executes and returns the result of given function, if given key is duplicated,
// it will wait the first called result and don't execute.
//
// 执行给定函数并返回数据, 如果给定 key 重复, 将等待第一次执行结果且不会重复执行.
func (s *Singleflight[K, V]) Do(key K, fn func() (V, error)) (V, error) {
	s.mu.Lock()
	if s.m == nil {
		s.m = make(map[K]*call[V])
	}

	// 如果相同 key 函数执行已存在, 则直接等待结果, 不重复执行
	if c, ok := s.m[key]; ok {
		s.mu.Unlock()
		c.wg.Wait()
		return c.value, c.err
	}

	// 如果不存在, 则创建一次函数执行
	c := new(call[V])
	c.wg.Add(1)
	s.m[key] = c
	s.mu.Unlock()

	// 执行函数
	s.doCall(c, key, fn)
	return c.value, c.err
}

func (s *Singleflight[K, V]) doCall(c *call[V], key K, fn func() (V, error)) {
	// 函数执行完之后删除记录
	defer func() {
		s.mu.Lock()
		defer s.mu.Unlock()
		c.wg.Done()
		if s.m[key] == c {
			delete(s.m, key)
		}
	}()

	// 捕获函数执行的 panic
	defer func() {
		if x := recover(); x != nil {
			c.err = fmt.Errorf("panic: %v\n\n%s", x, debug.Stack())
		}
	}()

	c.value, c.err = fn()
}
