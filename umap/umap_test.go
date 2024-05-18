package umap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	assert.ElementsMatch(t, []int{1, 2, 3}, Keys(map[int]string{1: "1", 2: "2", 3: "3"}))
	assert.ElementsMatch(t, []string{"1", "2", "3"}, Keys(map[string]int{"1": 1, "2": 2, "3": 3}))
	assert.ElementsMatch(t, []int{1, 3, 2}, Keys(map[int]int{1: 1, 2: 2, 3: 3}))
}

func TestValues(t *testing.T) {
	assert.ElementsMatch(t, []string{"1", "2", "3"}, Values(map[int]string{1: "1", 2: "2", 3: "3"}))
	assert.ElementsMatch(t, []int{1, 2, 3}, Values(map[string]int{"1": 1, "2": 2, "3": 3}))
	assert.ElementsMatch(t, []int{1, 3, 2}, Values(map[int]int{1: 1, 2: 2, 3: 3}))
	assert.ElementsMatch(t, []int{1, 3, 3, 2}, Values(map[int]int{1: 1, 2: 2, 3: 3, 100: 3}))
}
