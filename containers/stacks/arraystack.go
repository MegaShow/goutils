package stacks

// ArrayStack is a stack implemented by array.
//
// 一个由数组实现的栈.
type ArrayStack[E any] struct {
	items []E
}

// NewArrayStack create an array stack.
//
// 创建一个数组栈.
func NewArrayStack[E any]() *ArrayStack[E] {
	return &ArrayStack[E]{}
}

// Clear removes all elements of stack.
//
// 移除栈的所有元素.
func (s *ArrayStack[E]) Clear() {
	s.items = s.items[:0]
}

// Len returns length of stack.
//
// 返回栈长度.
func (s *ArrayStack[E]) Len() int {
	return len(s.items)
}

// Peek try to get the top element of stack.
// The second return value will be false if stack is empty.
//
// 获取栈顶元素. 如果栈为空, 则第二个返回值为 false.
func (s *ArrayStack[E]) Peek() (E, bool) {
	if len(s.items) == 0 {
		var zero E
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// Pop removes the top element of stack.
// The second return value will be false if stack is empty.
//
// 移除栈顶元素. 如果栈为空, 则第二个返回值为 false.
func (s *ArrayStack[E]) Pop() (E, bool) {
	if len(s.items) == 0 {
		var zero E
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

// Push inserts a new element at the top of stack.
//
// 插入新元素到栈顶.
func (s *ArrayStack[E]) Push(item E) {
	s.items = append(s.items, item)
}
