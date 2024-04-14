package uslice

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	assert.Equal(t, []int{2, 3}, Filter([]int{1, 2, 3}, func(v int) bool { return v > 1 }))
	assert.Equal(t, []string{"1", "2", "3"}, Filter([]string{"1", "2", "3"}, func(v string) bool { return v != "" }))
}

func TestMap(t *testing.T) {
	assert.Equal(t, []int{2, 3, 4}, Map([]int{1, 2, 3}, func(v int) int { return v + 1 }))
	assert.Equal(t, []string{"1", "2", "3"}, Map([]int{1, 2, 3}, func(v int) string { return strconv.FormatInt(int64(v), 10) }))
}
