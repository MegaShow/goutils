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
	clear(items[i:]) // for GC
	return items[:i]
}

// Find returns the first item which satisfy the provided function.
//
// 返回第一个满足给定函数条件的元素.
func Find[T any](items []T, fn func(item T) bool) (item T, ok bool) {
	for _, item := range items {
		if fn(item) {
			return item, true
		}
	}
	return item, false
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
	clear(items[i:]) // for GC
	return items[:i]
}

// GroupBy coverts slice to map and group by key.
//
// 将切片转换成 map, 并根据 key 分组.
func GroupBy[T any, K comparable](items []T, conv func(item T) K) map[K][]T {
	results := make(map[K][]T, len(items))
	for _, item := range items {
		key := conv(item)
		results[key] = append(results[key], item)
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

// Of creates a slice with items, if no items returns empty slice not nil.
//
// 创建一个切片, 如果未给定元素则返回空切片.
func Of[T any](items ...T) []T {
	return items
}

// ToMap converts slice to map.
//
// 将切片转换成 map.
func ToMap[T any, K comparable](items []T, conv func(item T) K) map[K]T {
	results := make(map[K]T, len(items))
	for _, item := range items {
		results[conv(item)] = item
	}
	return results
}
