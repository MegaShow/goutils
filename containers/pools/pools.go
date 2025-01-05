package pools

import (
	"sync"

	"go.icytown.com/utils/ucond"
)

// Pool is a set of temporary objects that can be reused.
//
// 一个可以被重复使用的临时对象的集合.
type Pool[T any] struct {
	pool sync.Pool
}

// NewPool creates a pool instance.
//
// 创建一个池实例.
func NewPool[T any](newFunc func() T) *Pool[T] {
	return &Pool[T]{
		pool: sync.Pool{
			New: ucond.If(newFunc != nil, func() any { return newFunc() }, nil),
		},
	}
}

// Get returns an arbitrary item from the pool, and removes it from the pool.
// If no item can be selected from the pool, Get will returns the result of calling newFunc.
//
// 从池中取出任意一个元素并返回. 如果没有元素能从池中被选择, 则返回调用 newFunc 的结果.
func (p *Pool[T]) Get() T {
	x, _ := p.pool.Get().(T)
	return x
}

// Put adds x to the pool.
//
// 将 x 添加到池中.
func (p *Pool[T]) Put(x T) {
	p.pool.Put(x)
}
