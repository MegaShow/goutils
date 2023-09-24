package stacks

import "testing"

func TestArrayStack_Clear(t *testing.T) {
	s := NewArrayStack[int]()
	s.Push(1)
	s.Clear()
	if got := s.Len(); got != 0 {
		t.Errorf("Len(): %d, want: 0", got)
	}
}

func TestArrayStack_Len(t *testing.T) {
	s := NewArrayStack[int]()
	if got := s.Len(); got != 0 {
		t.Errorf("Len(): %d, want: 0", got)
	}
	s.Push(1)
	if got := s.Len(); got != 1 {
		t.Errorf("Len(): %d, want: 1", got)
	}
	s.Push(2)
	s.Push(3)
	if got := s.Len(); got != 3 {
		t.Errorf("Len(): %d, want: 3", got)
	}
}
