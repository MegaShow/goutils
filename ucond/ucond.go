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

// Must panics if err is not nil. If err is nil, returns value.
//
// err 不为 nil 时主动 panic, err 为 nil 时返回 value.
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

// Must0 has same behavior as Must, but returns 0 variables.
//
// 与 Must 行为一致, 但不返回值.
func Must0(err error) {
	if err != nil {
		panic(err)
	}
}

// Must1 has same behavior as Must.
//
// 与 Must 行为一致.
func Must1[T1 any](v1 T1, err error) T1 {
	return Must(v1, err)
}

// Must2 has same behavior as Must, but returns 2 variables.
//
// 与 Must 行为一致, 但返回 2 个值.
func Must2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return v1, v2
}

// Must3 has same behavior as Must, but returns 3 variables.
//
// 与 Must 行为一致, 但返回 3 个值.
func Must3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	if err != nil {
		panic(err)
	}
	return v1, v2, v3
}

// Must4 has same behavior as Must, but returns 4 variables.
//
// 与 Must 行为一致, 但返回 4 个值.
func Must4[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4, err error) (T1, T2, T3, T4) {
	if err != nil {
		panic(err)
	}
	return v1, v2, v3, v4
}

// Not creates a new function that returns opposite bool value.
//
// 创建一个返回相反布尔值的函数.
func Not[T any](fn func(T) bool) func(T) bool {
	return func(value T) bool {
		return !fn(value)
	}
}
