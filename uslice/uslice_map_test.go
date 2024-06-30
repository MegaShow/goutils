package uslice

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupBy(t *testing.T) {
	assert.Equal(t, map[int][]int{1: {1, 5, 5, 3}, 0: {4, 6}}, GroupBy([]int{1, 4, 5, 5, 3, 6}, func(v int) int { return v % 2 }))
	assert.Equal(t, map[bool][]int{false: {1, 2, 2, 3}, true: {4, 6}}, GroupBy([]int{1, 4, 2, 2, 3, 6}, func(v int) bool { return v > 3 }))
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
