package stacks

import "encoding/json"

var _ Stack[int] = (*ArrayStack[int])(nil)

// ArrayStack is a stack implemented by array.
//
// 一个由数组实现的栈.
type ArrayStack[T any] struct {
	items []T
}

// NewArrayStack create an array stack.
//
// 创建一个数组栈.
func NewArrayStack[T any](opts ...Option) *ArrayStack[T] {
	s := new(ArrayStack[T])

	for _, opt := range opts {
		if opt.capacity != 0 {
			s.items = make([]T, 0, opt.capacity)
		}
	}

	return s
}

// Peek try to get the top item of stack.
// The second return value will be false if stack is empty.
//
// 获取栈顶元素. 如果栈为空, 则第二个返回值为 false.
func (s *ArrayStack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// Pop removes the top item of stack.
// The second return value will be false if stack is empty.
//
// 移除栈顶元素. 如果栈为空, 则第二个返回值为 false.
func (s *ArrayStack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

// Push inserts a new item at the top of stack.
//
// 插入新元素到栈顶.
func (s *ArrayStack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Clear removes all items of stack.
//
// 移除栈的所有元素.
func (s *ArrayStack[T]) Clear() {
	s.items = s.items[:0]
}

// Len returns length of stack.
//
// 返回栈长度.
func (s *ArrayStack[T]) Len() int {
	return len(s.items)
}

// Values returns all items as slice.
//
// 返回包含所有元素的切片.
func (s *ArrayStack[T]) Values() []T {
	items := make([]T, len(s.items))
	copy(items, s.items)
	return items
}

// MarshalJSON marshal stack into json bytes.
//
// 序列化集合.
func (s *ArrayStack[T]) MarshalJSON() ([]byte, error) {
	values := s.Values()
	return json.Marshal(values)
}

// UnmarshalJSON unmarshal json bytes to stack.
//
// 反序列化集合.
func (s *ArrayStack[T]) UnmarshalJSON(data []byte) error {
	values := make([]T, 0)
	err := json.Unmarshal(data, &values)
	if err != nil {
		return err
	}
	s.items = values
	return nil
}
