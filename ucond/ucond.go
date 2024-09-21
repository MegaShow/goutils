// Package ucond provides utils of condition.
//
// 包 ucond 提供了条件相关的工具.
package ucond

// If returns trueValue if cond is true, if cond is false then returns falseValue.
//
// 根据 cond 的值决定返回 trueValue 或 falseValue.
func If[T any](cond bool, trueValue, falseValue T) T {
	if cond {
		return trueValue
	}
	return falseValue
}

// If runs trueFn if cond is true, if cond is false then runs falseFn.
//
// 根据 cond 的值决定执行返回 trueFn 或 falseFn 执行后的返回值.
func IfFunc[T any](cond bool, trueFn, falseFn func() T) T {
	if cond {
		return trueFn()
	}
	return falseFn()
}

// Not creates a new function that returns opposite bool value.
//
// 创建一个返回相反布尔值的函数.
func Not[T any](fn func(T) bool) func(T) bool {
	return func(value T) bool {
		return !fn(value)
	}
}
