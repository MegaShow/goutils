package ucond

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIf(t *testing.T) {
	assert.Equal(t, true, If(true, true, false))
	assert.Equal(t, 1, If(true, 1, 2))
	assert.Equal(t, 2, If(false, 1, 2))
}

func TestIfFunc(t *testing.T) {
	assert.Equal(t, true, IfFunc(true, func() bool { return true }, func() bool { return false }))
	assert.Equal(t, 1, IfFunc(true, func() int { return 1 }, func() int { return 2 }))
	assert.Equal(t, 2, IfFunc(false, func() int { return 1 }, func() int { return 2 }))
}
