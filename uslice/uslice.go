// Package uslice provides utils of slice.
//
// 包 uslice 提供了切片相关的工具.
package uslice

// Filter create a new slice populated with the items that satisfy the provided function.
//
// 创建一个新的数组, 这个数组由满足给定函数的元素组成.
func Filter[T any](items []T, fn func(item T) bool) []T {
	results := make([]T, 0, len(items))
	for _, item := range items {
		if fn(item) {
			results = append(results, item)
		}
	}
	return results
}

// Map creates a new slice populated with the results of calling a provided function on every item in the calling array.
//
// 创建一个新的数组, 这个新数组由原数组中的每个元素调用一次提供的函数后的返回值组成.
func Map[T, R any](items []T, fn func(item T) R) []R {
	results := make([]R, len(items))
	for idx, item := range items {
		results[idx] = fn(item)
	}
	return results
}
