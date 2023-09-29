// Package stacks implements stacks of different type, such as [ArrayStack].
//
// stacks 包实现了多种栈数据结构, 比如 [ArrayStack].
package stacks

import "icytown.com/goutils/containers/internal"

// Stack is the interface for stacks.
//
// 栈接口.
type Stack[T any] interface {
	Peek() (T, bool)
	Pop() (T, bool)
	Push(item T)

	internal.Container[T]
}
