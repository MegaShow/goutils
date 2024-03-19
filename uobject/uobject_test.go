package uobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	assert.Equal(t, 12, Default(12, 100))
	assert.Equal(t, 100, Default(0, 100))
	assert.Equal(t, 12.0, Default(12.0, 100))
	assert.Equal(t, 100.0, Default(0.0, 100))
	assert.Equal(t, "hello", Default("hello", "default"))
	assert.Equal(t, "default", Default("", "default"))
}

func TestIndirect(t *testing.T) {
	var v1, v2 int = 0, 12
	assert.Equal(t, 0, Indirect(&v1))
	assert.Equal(t, 12, Indirect(&v2))
	assert.Equal(t, 0, Indirect[int](nil))
}

func TestIndirectOr(t *testing.T) {
	var v1, v2 int = 0, 12
	assert.Equal(t, 0, IndirectOr(&v1, 12))
	assert.Equal(t, 12, IndirectOr(&v2, 13))
	assert.Equal(t, 0, IndirectOr(nil, 0))
	assert.Equal(t, 12, IndirectOr(nil, 12))
}

func TestPtr(t *testing.T) {
	var v1, v2 int = 0, 12
	assert.Equal(t, &v1, Ptr(0))
	assert.Equal(t, &v2, Ptr(12))
}
