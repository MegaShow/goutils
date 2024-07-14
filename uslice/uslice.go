// Package uslice provides utils of slice.
//
// 包 uslice 提供了切片相关的工具.
package uslice

// Distinct modify source slice and returns modified slice with distinct items.
//
// 修改原来的数组, 确保数组每一个元素只存在一次, 并返回修改后的数组.
func Distinct[T comparable](items []T) []T {
	exists := make(map[T]struct{}, len(items))
	i := 0
	for k := 0; k < len(items); k++ {
		_, ok := exists[items[k]]
		if !ok {
			if i != k {
				items[i] = items[k]
			}
			i++
			exists[items[k]] = struct{}{}
		}
	}
	clear(items[i:])
	return items[:i]
}

// Filter modify source slice and returns modified slice with items that satisfy the provided function.
//
// 修改原来的数组, 确保数组每一个元素都满足给定函数条件, 并返回修改后的数组.
func Filter[T any](items []T, fn func(item T) bool) []T {
	i := 0
	for k := 0; k < len(items); k++ {
		if fn(items[k]) {
			if i != k {
				items[i] = items[k]
			}
			i++
		}
	}
	clear(items[i:])
	return items[:i]
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
