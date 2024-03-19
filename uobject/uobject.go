// Package uobject provides utils of variable.
//
// 包 uobject 提供了变量相关的工具.
package uobject

// Default returns a default value if value is zero.
//
// 返回默认值如果给定的变量为零值.
func Default[T comparable](value, defaultValue T) T {
	var zero T
	if value == zero {
		return defaultValue
	}
	return value
}

// Indirect returns the value that ptr points to. If ptr is nil, Indirect returns a zero value.
//
// 返回指针指向的值, 如果指针为 nil, 则返回空值.
func Indirect[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}

// IndirectOr returns the value that ptr points to. If ptr is nil, Indirect returns a default value.
//
// 返回指针指向的值, 如果指针为 nil, 则返回传入的默认值.
func IndirectOr[T any](valuePtr *T, defaultValue T) T {
	if valuePtr == nil {
		return defaultValue
	}
	return *valuePtr
}

// Ptr returns a pointer that has same value.
//
// 返回指针变量, 其指向的值与传入的值相同.
func Ptr[T any](value T) *T {
	return &value
}
