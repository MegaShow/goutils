// Package umath provides utils of math.
//
// 包 umath 提供了数学相关的工具.
package umath

import "math"

// FloorFloat returns the least float64 value greater than or equal to x with special precision.
//
// 返回一个最小的比给定值大或相等的, 具备特定精度的浮点数值.
func CeilFloat(x float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Ceil(x*p) / p
}

// FloorFloat returns the greastest float64 value less than or equal to x with special precision.
//
// 返回一个最大的比给定值小或相等的, 具备特定精度的浮点数值.
func FloorFloat(x float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(x*p) / p
}

// RoundFloat returns a nearest float64 value with special precision.
//
// 返回一个与给定值最接近的具备特定精度的浮点数值.
func RoundFloat(x float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Round(x*p) / p
}
