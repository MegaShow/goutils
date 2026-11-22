// Package umath provides utils of math.
//
// 包 umath 提供了数学相关的工具.
package umath

import "math"

// CeilFloat rounds x up to the given decimal precision.
// Positive precision keeps fractional digits; negative precision rounds to powers of ten.
// It uses float64 scaling; extreme values or precision may cause overflow, underflow, Inf or NaN.
//
// 按指定十进制精度向上取整, 正精度保留小数位, 负精度取整到十的幂次位
// 使用 float64 缩放, 极端数值或精度可能产生溢出、下溢、Inf 或 NaN
func CeilFloat(x float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Ceil(x*p) / p
}

// FloorFloat rounds x down to the given decimal precision.
// Positive precision keeps fractional digits; negative precision rounds to powers of ten.
// It uses float64 scaling; extreme values or precision may cause overflow, underflow, Inf or NaN.
//
// 按指定十进制精度向下取整, 正精度保留小数位, 负精度取整到十的幂次位
// 使用 float64 缩放, 极端数值或精度可能产生溢出、下溢、Inf 或 NaN
func FloorFloat(x float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(x*p) / p
}

// RoundFloat rounds x to the nearest value at the given decimal precision, with ties away from zero.
// Positive precision keeps fractional digits; negative precision rounds to powers of ten.
// It uses float64 scaling; extreme values or precision may cause overflow, underflow, Inf or NaN.
//
// 按指定十进制精度四舍五入, 中间值向远离零的方向取整
// 正精度保留小数位, 负精度取整到十的幂次位
// 使用 float64 缩放, 极端数值或精度可能产生溢出、下溢、Inf 或 NaN
func RoundFloat(x float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Round(x*p) / p
}
