package uslice

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
