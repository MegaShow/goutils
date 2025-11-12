package uobject

import (
	"testing"

	"go.icytown.com/utils/internal/assert"
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
	var v1, v2 = 0, 12
	assert.Equal(t, 0, Indirect(&v1))
	assert.Equal(t, 12, Indirect(&v2))
	assert.Equal(t, 0, Indirect[int](nil))
}

func TestIndirectOr(t *testing.T) {
	var v1, v2 = 0, 12
	assert.Equal(t, 0, IndirectOr(&v1, 12))
	assert.Equal(t, 12, IndirectOr(&v2, 13))
	assert.Equal(t, 0, IndirectOr(nil, 0))
	assert.Equal(t, 12, IndirectOr(nil, 12))
}

func TestIsNotZero(t *testing.T) {
	assert.True(t, IsNotZero(1))
	assert.False(t, IsNotZero(0))
	assert.True(t, IsNotZero("1"))
	assert.False(t, IsNotZero(""))
}

func TestIsZero(t *testing.T) {
	assert.False(t, IsZero(1))
	assert.True(t, IsZero(0))
	assert.False(t, IsZero("1"))
	assert.True(t, IsZero(""))
}

func TestPtr(t *testing.T) {
	var v1, v2 = 0, 12
	assert.Equal(t, &v1, Ptr(0))
	assert.Equal(t, &v2, Ptr(12))
}
