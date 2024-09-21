package uslice

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	value, ok := Find([]int{1, 2, 3}, func(v int) bool { return v == 1 })
	assert.Equal(t, 1, value)
	assert.True(t, ok)
	value, ok = Find([]int{1, 2, 3}, func(v int) bool { return v > 0 })
	assert.Equal(t, 1, value)
	assert.True(t, ok)
	value, ok = Find([]int{1, 2, 3}, func(v int) bool { return v < 0 })
	assert.Equal(t, 0, value)
	assert.False(t, ok)
}

func TestFilter(t *testing.T) {
	assert.Equal(t, []int{2, 3}, Filter([]int{1, 2, 3}, func(v int) bool { return v > 1 }))
	assert.Equal(t, []string{"1", "2", "3"}, Filter([]string{"1", "2", "3"}, func(v string) bool { return v != "" }))
}

func TestGroupBy(t *testing.T) {
	assert.Equal(t, map[int][]int{1: {1, 5, 5, 3}, 0: {4, 6}}, GroupBy([]int{1, 4, 5, 5, 3, 6}, func(v int) int { return v % 2 }))
	assert.Equal(t, map[bool][]int{false: {1, 2, 2, 3}, true: {4, 6}}, GroupBy([]int{1, 4, 2, 2, 3, 6}, func(v int) bool { return v > 3 }))
}

func TestMap(t *testing.T) {
	assert.Equal(t, []int{2, 3, 4}, Map([]int{1, 2, 3}, func(v int) int { return v + 1 }))
	assert.Equal(t, []string{"1", "2", "3"}, Map([]int{1, 2, 3}, func(v int) string { return strconv.FormatInt(int64(v), 10) }))
}

func TestToMap(t *testing.T) {
	assert.Equal(t, map[string]int{"1": 1, "2": 2}, ToMap([]int{1, 2}, func(v int) string { return strconv.FormatInt(int64(v), 10) }))
	type obj struct {
		id   int
		name string
	}
	assert.Equal(t, map[int]obj{1: {id: 1, name: "name-1"}, 2: {id: 2, name: "name-2"}},
		ToMap([]obj{{id: 1, name: "name-1"}, {id: 2, name: "name-2"}}, func(v obj) int { return v.id }))
}

func TestUnique(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5}, Unique([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, []int{1, 2, 4}, Unique([]int{1, 2, 2, 4, 4}))
	assert.Equal(t, []int{1, 2, 4}, Unique([]int{1, 2, 2, 4, 4, 1}))
	assert.Equal(t, []int{1, 2, 3}, Unique([]int{1, 2, 2, 3, 1, 1}))
}

func TestUniqueFunc(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5}, UniqueFunc([]int{1, 2, 3, 4, 5}, func(v int) int { return v }))
	assert.Equal(t, []int{1, 2, 4}, UniqueFunc([]int{1, 2, 2, 4, 4}, func(v int) int { return v }))
	assert.Equal(t, []int{1, 2, 4}, UniqueFunc([]int{1, 2, 2, 4, 4, 1}, func(v int) int { return v }))
	assert.Equal(t, []int{1, 2, 3}, UniqueFunc([]int{1, 2, 2, 3, 1, 1}, func(v int) int { return v }))
}
