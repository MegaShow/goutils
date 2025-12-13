// Package uslice provides utils of slice.
//
// 包 uslice 提供了切片相关的工具.
package uslice

// All returns true if all items satisfy the match function.
//
// 返回 true 如果所有元素都满足给定函数条件.
func All[T any](items []T, match func(item T) bool) bool {
	for _, item := range items {
		if !match(item) {
			return false
		}
	}
	return true
}

// Find returns the first item which satisfy the match function.
// If wants to get index please use standard [slices.IndexFunc].
//
// 返回第一个满足给定函数条件的元素.
// 如果期望获取元素下标请使用标准库函数 [slices.IndexFunc].
func Find[T any](items []T, match func(item T) bool) (res T, ok bool) {
	for _, item := range items {
		if match(item) {
			return item, true
		}
	}
	return
}

// Find returns the last item which satisfy the match function.
//
// 返回最后一个满足给定函数条件的元素.
func FindLast[T any](items []T, match func(item T) bool) (res T, ok bool) {
	for i := len(items) - 1; i >= 0; i-- {
		if match(items[i]) {
			return items[i], true
		}
	}
	return
}

// Filter creates a new slice with items that satisfy the match function.
// If wants to modify source slice, please use [slices.DeleteFunc].
//
// 创建一个新的切片, 确保切片每一个元素都满足给定函数条件.
// 如果想要在原切片上修改, 请使用 [slices.DeleteFunc].
func Filter[T any](items []T, match func(item T) bool) []T {
	results := make([]T, 0, len(items))
	for _, item := range items {
		if match(item) {
			results = append(results, item)
		}
	}
	return results
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
// 创建一个新的切片, 这个新切片由原切片中的每个元素调用一次提供的函数后的返回值组成.
func Map[T, R any](items []T, conv func(item T) R) []R {
	results := make([]R, len(items))
	for idx, item := range items {
		results[idx] = conv(item)
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

// Unique creates a new slice with unique items.
// If wants to modify source slice, please use [slices.Compact] for sorted slice.
//
// 创建一个新的切片, 确保切片每一个元素只存在一次.
// 如果想要在原切片上修改, 请对有序的切片使用 [slices.Compact].
func Unique[T comparable](items []T) []T {
	exists := make(map[T]struct{}, len(items))
	results := make([]T, 0, len(items))
	for _, item := range items {
		if _, ok := exists[item]; ok {
			continue
		}
		results = append(results, item)
		exists[item] = struct{}{}
	}
	return results
}

// UniqueFunc creates a new slice with unique key of provided function.
// If wants to modify source slice, please use [slices.CompactFunc] for sorted slice.
//
// 创建一个新的切片, 确保切片每一个元素对应函数生成的键只存在一次.
// 如果想要在原切片上修改, 请对有序的切片使用 [slices.CompactFunc].
func UniqueFunc[T any, K comparable](items []T, hash func(item T) K) []T {
	exists := make(map[K]struct{}, len(items))
	results := make([]T, 0, len(items))
	for _, item := range items {
		k := hash(item)
		if _, ok := exists[k]; ok {
			continue
		}
		results = append(results, item)
		exists[k] = struct{}{}
	}
	return results
}
