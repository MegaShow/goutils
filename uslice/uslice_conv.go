package uslice

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
