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
