package stacks

import (
	"encoding/json"
	"testing"

	"go.icytown.com/utils/internal/assert"
)

func TestArrayStack_Clear(t *testing.T) {
	t.Parallel()
	s := NewArrayStack[int]()
	s.Push(1)
	s.Clear()
	assert.Equal(t, 0, s.Len())

	// 测试清空后释放引用并保留容量
	pointerStack := NewArrayStack[*int](WithCapacity(4))
	first, second := 1, 2
	pointerStack.Push(&first)
	pointerStack.Push(&second)
	pointerStack.Clear()
	assert.Equal(t, 4, cap(pointerStack.items))
	assert.Equal(t, 0, pointerStack.Len())
	for _, item := range pointerStack.items[:cap(pointerStack.items)] {
		assert.Nil(t, item)
	}
}

func TestArrayStack_Len(t *testing.T) {
	t.Parallel()
	s := NewArrayStack[int]()
	assert.Equal(t, 0, s.Len())
	s.Push(1)
	assert.Equal(t, 1, s.Len())
	s.Push(2)
	s.Push(3)
	assert.Equal(t, 3, s.Len())
}

func TestArrayStack_MarshalJSON(t *testing.T) {
	t.Parallel()
	s := NewArrayStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	bytes, err := json.Marshal(s)
	assert.Nil(t, err)
	assert.Equal(t, "[1,2,3]", string(bytes))
}

func TestArrayStack_Peek(t *testing.T) {
	t.Parallel()
	s := NewArrayStack[int]()
	top, ok := s.Peek()
	assert.False(t, ok)
	assert.Zero(t, top)
	s.Push(1)
	top, ok = s.Peek()
	assert.True(t, ok)
	assert.Equal(t, 1, top)
	assert.Equal(t, 1, s.Len())
}

func TestArrayStack_Pop(t *testing.T) {
	t.Parallel()
	s := NewArrayStack[int]()
	top, ok := s.Pop()
	assert.False(t, ok)
	assert.Zero(t, top)
	s.Push(1)
	top, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, top)
	assert.Equal(t, 0, s.Len())

	// 测试弹出后释放底层数组对元素的引用
	pointerStack := NewArrayStack[*int]()
	item := 42
	pointerStack.Push(&item)
	pointer, ok := pointerStack.Pop()
	assert.True(t, ok)
	assert.Equal(t, &item, pointer)
	assert.Nil(t, pointerStack.items[:cap(pointerStack.items)][0])
}

func TestArrayStack_UnmarshalJSON(t *testing.T) {
	t.Parallel()
	var s *ArrayStack[int]
	err := json.Unmarshal([]byte("[1,2,3]"), &s)
	assert.Nil(t, err)
	assert.Equal(t, []int{1, 2, 3}, s.Values())

	// 测试失败时保留原栈
	err = json.Unmarshal([]byte(`[4,"invalid"]`), s)
	assert.NotNil(t, err)
	assert.Equal(t, []int{1, 2, 3}, s.Values())
}

func TestArrayStack_Values(t *testing.T) {
	t.Parallel()
	s := NewArrayStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.Equal(t, []int{1, 2, 3}, s.Values())
	s.Push(0)
	assert.Equal(t, []int{1, 2, 3, 0}, s.Values())
}

func TestNewArrayStack(t *testing.T) {
	t.Parallel()
	s := NewArrayStack[int]()
	assert.NotNil(t, s)
	assert.Equal(t, 0, cap(s.items))

	s = NewArrayStack[int](WithCapacity(10))
	assert.NotNil(t, s)
	assert.Equal(t, 10, cap(s.items))
}
