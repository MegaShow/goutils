package umath

import (
	"math"
	"testing"

	"go.icytown.com/utils/internal/assert"
)

func TestCeilFloat(t *testing.T) {
	t.Parallel()
	assert.Equal(t, float64(124), CeilFloat(123.456789, 0))
	assert.Equal(t, float64(124), CeilFloat(123.56789, 0))
	assert.Equal(t, float64(123.5), CeilFloat(123.456789, 1))
	assert.Equal(t, float64(123.46), CeilFloat(123.456789, 2))

	// 测试负精度和特殊浮点值
	assert.Equal(t, float64(130), CeilFloat(123, -1))
	assert.Equal(t, float64(-120), CeilFloat(-123, -1))
	assert.Equal(t, float64(0), CeilFloat(0, 308))
	assert.True(t, math.Signbit(CeilFloat(math.Copysign(0, -1), 2)))
	assert.True(t, math.IsNaN(CeilFloat(math.NaN(), 2)))
	assert.True(t, math.IsInf(CeilFloat(math.Inf(1), 2), 1))
	assert.True(t, math.IsInf(CeilFloat(math.Inf(-1), 2), -1))
}

func TestFloorFloat(t *testing.T) {
	t.Parallel()
	assert.Equal(t, float64(123), FloorFloat(123.456789, 0))
	assert.Equal(t, float64(123), FloorFloat(123.56789, 0))
	assert.Equal(t, float64(123.4), FloorFloat(123.456789, 1))
	assert.Equal(t, float64(123.45), FloorFloat(123.456789, 2))

	// 测试负精度和特殊浮点值
	assert.Equal(t, float64(120), FloorFloat(123, -1))
	assert.Equal(t, float64(-130), FloorFloat(-123, -1))
	assert.Equal(t, float64(0), FloorFloat(0, 308))
	assert.True(t, math.Signbit(FloorFloat(math.Copysign(0, -1), 2)))
	assert.True(t, math.IsNaN(FloorFloat(math.NaN(), 2)))
	assert.True(t, math.IsInf(FloorFloat(math.Inf(1), 2), 1))
	assert.True(t, math.IsInf(FloorFloat(math.Inf(-1), 2), -1))
}

func TestRoundFloat(t *testing.T) {
	t.Parallel()
	assert.Equal(t, float64(123), RoundFloat(123.456789, 0))
	assert.Equal(t, float64(124), RoundFloat(123.56789, 0))
	assert.Equal(t, float64(123.5), RoundFloat(123.456789, 1))
	assert.Equal(t, float64(123.46), RoundFloat(123.456789, 2))

	// 测试负精度和特殊浮点值
	assert.Equal(t, float64(120), RoundFloat(123, -1))
	assert.Equal(t, float64(-120), RoundFloat(-123, -1))
	assert.Equal(t, float64(0), RoundFloat(0, 308))
	assert.True(t, math.Signbit(RoundFloat(math.Copysign(0, -1), 2)))
	assert.True(t, math.IsNaN(RoundFloat(math.NaN(), 2)))
	assert.True(t, math.IsInf(RoundFloat(math.Inf(1), 2), 1))
	assert.True(t, math.IsInf(RoundFloat(math.Inf(-1), 2), -1))
}
