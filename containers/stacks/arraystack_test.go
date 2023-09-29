package stacks

import (
	"encoding/json"
	"testing"

	"icytown.com/goutils/testings/assert"
)

func TestArrayStack_Peek(t *testing.T) {
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
	s := NewArrayStack[int]()
	top, ok := s.Pop()
	assert.False(t, ok)
	assert.Zero(t, top)
	s.Push(1)
	top, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, top)
	assert.Equal(t, 0, s.Len())
}

func TestArrayStack_Clear(t *testing.T) {
	s := NewArrayStack[int]()
	s.Push(1)
	s.Clear()
	assert.Equal(t, 0, s.Len())
}

func TestArrayStack_Len(t *testing.T) {
	s := NewArrayStack[int]()
	assert.Equal(t, 0, s.Len())
	s.Push(1)
	assert.Equal(t, 1, s.Len())
	s.Push(2)
	s.Push(3)
	assert.Equal(t, 3, s.Len())
}

func TestArrayStack_Values(t *testing.T) {
	s := NewArrayStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.Equal(t, []int{1, 2, 3}, s.Values())
	s.Push(0)
	assert.Equal(t, []int{1, 2, 3, 0}, s.Values())
}

func TestArrayStack_MarshalJSON(t *testing.T) {
	s := NewArrayStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	bytes, err := json.Marshal(s)
	assert.Nil(t, err)
	assert.Equal(t, "[1,2,3]", string(bytes))
}

func TestArrayStack_UnmarshalJSON(t *testing.T) {
	var s *ArrayStack[int]
	err := json.Unmarshal([]byte("[1,2,3]"), &s)
	assert.Nil(t, err)
	assert.Equal(t, []int{1, 2, 3}, s.Values())
}
