package sets

import (
	"encoding/json"
	"testing"

	"go.icytown.com/utils/internal/assert"
)

func TestNewHashSet(t *testing.T) {
	s := NewHashSet[int]()
	assert.NotNil(t, s)
	s = NewHashSet[int](WithCapacity(10))
	assert.NotNil(t, s)
}

func TestHashSet_Add(t *testing.T) {
	s := NewHashSet[int]()
	s.Add(1)
	assert.True(t, s.Contains(1))
	s.Add(2)
	assert.True(t, s.Contains(2))
	s.Add(2)
	assert.Equal(t, 2, s.Len())
}

func TestHashSet_Contains(t *testing.T) {
	s := NewHashSet[int]()
	s.Add(1)
	s.Add(2)
	assert.False(t, s.Contains(0))
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.False(t, s.Contains(3))
}

func TestHashSet_Remove(t *testing.T) {
	s := NewHashSet[int]()
	s.Add(1)
	s.Add(2)
	assert.False(t, s.Remove(0))
	if assert.True(t, s.Remove(1)) {
		assert.False(t, s.Contains(1))
	}
}

func TestHashSet_Clear(t *testing.T) {
	s := NewHashSet[int]()
	s.Add(1)
	s.Clear()
	assert.Equal(t, 0, s.Len())
}

func TestHashSet_Len(t *testing.T) {
	s := NewHashSet[int]()
	assert.Equal(t, 0, s.Len())
	s.Add(1)
	assert.Equal(t, 1, s.Len())
	s.Add(2)
	s.Add(3)
	assert.Equal(t, 3, s.Len())
	s.Add(2)
	assert.Equal(t, 3, s.Len())
}

func TestHashSet_Values(t *testing.T) {
	s := NewHashSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	assert.ElementsMatch(t, []int{1, 2, 3}, s.Values())
	s.Add(0)
	assert.ElementsMatch(t, []int{0, 1, 2, 3}, s.Values())
}

func TestHashSet_MarshalJSON(t *testing.T) {
	s := NewHashSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	setBytes, err := json.Marshal(s)
	assert.Nil(t, err)
	items := make([]int, 0)
	err = json.Unmarshal(setBytes, &items)
	assert.Nil(t, err)
	assert.ElementsMatch(t, []int{1, 2, 3}, items)
}

func TestHashSet_UnmarshalJSON(t *testing.T) {
	var s *HashSet[int]
	err := json.Unmarshal([]byte("[1,2,3]"), &s)
	assert.Nil(t, err)
	assert.ElementsMatch(t, []int{1, 2, 3}, s.Values())
}
