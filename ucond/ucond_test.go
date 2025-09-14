package ucond

import (
	"errors"
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

func TestMust(t *testing.T) {
	assert.NotPanics(t, func() { assert.Equal(t, 1, Must(1, nil)) })
	assert.NotPanics(t, func() { assert.Equal(t, "test", Must("test", nil)) })
	assert.Panics(t, func() { Must(0, errors.New("test error")) })
}

func TestMust0(t *testing.T) {
	assert.NotPanics(t, func() { Must0(nil) })
	assert.Panics(t, func() { Must0(errors.New("test error")) })
}

func TestMust1(t *testing.T) {
	assert.NotPanics(t, func() { assert.Equal(t, 1, Must1(1, nil)) })
	assert.NotPanics(t, func() { assert.Equal(t, "test", Must1("test", nil)) })
	assert.Panics(t, func() { Must1(0, errors.New("test error")) })
}

func TestMust2(t *testing.T) {
	assert.NotPanics(t, func() {
		v1, v2 := Must2(1, "test", nil)
		assert.Equal(t, 1, v1)
		assert.Equal(t, "test", v2)
	})
	assert.Panics(t, func() { Must2(0, "test", errors.New("test error")) })
}

func TestMust3(t *testing.T) {
	assert.NotPanics(t, func() {
		v1, v2, v3 := Must3(1, "test", 1.2, nil)
		assert.Equal(t, 1, v1)
		assert.Equal(t, "test", v2)
		assert.Equal(t, 1.2, v3)
	})
	assert.Panics(t, func() { Must3(0, "test", 1.2, errors.New("test error")) })
}

func TestMust4(t *testing.T) {
	assert.NotPanics(t, func() {
		v1, v2, v3, v4 := Must4(1, "test", 1.2, error(nil), nil)
		assert.Equal(t, 1, v1)
		assert.Equal(t, "test", v2)
		assert.Equal(t, 1.2, v3)
		assert.Nil(t, v4)
	})
	assert.Panics(t, func() {
		_, _, _, _ = Must4(0, "test", 1.2, error(nil), errors.New("test error"))
	})
}

func TestNot(t *testing.T) {
	isZero := func(v int) bool {
		return v == 0
	}
	assert.True(t, Not(isZero)(1))
	assert.False(t, Not(isZero)(0))
}
