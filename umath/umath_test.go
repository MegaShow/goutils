package umath

import (
	"testing"

	"go.icytown.com/utils/internal/assert"
)

func TestCeilFloat(t *testing.T) {
	assert.Equal(t, float64(124), CeilFloat(123.456789, 0))
	assert.Equal(t, float64(124), CeilFloat(123.56789, 0))
	assert.Equal(t, float64(123.5), CeilFloat(123.456789, 1))
	assert.Equal(t, float64(123.46), CeilFloat(123.456789, 2))
}

func TestFloorFloat(t *testing.T) {
	assert.Equal(t, float64(123), FloorFloat(123.456789, 0))
	assert.Equal(t, float64(123), FloorFloat(123.56789, 0))
	assert.Equal(t, float64(123.4), FloorFloat(123.456789, 1))
	assert.Equal(t, float64(123.45), FloorFloat(123.456789, 2))
}

func TestRoundFloat(t *testing.T) {
	assert.Equal(t, float64(123), RoundFloat(123.456789, 0))
	assert.Equal(t, float64(124), RoundFloat(123.56789, 0))
	assert.Equal(t, float64(123.5), RoundFloat(123.456789, 1))
	assert.Equal(t, float64(123.46), RoundFloat(123.456789, 2))
}
