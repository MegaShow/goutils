package umap

// Keys returns keys of map as a slice.
//
// 以切片形式返回 map 的所有键.
//
// Deprecated: As of Go 1.23, please use [maps.Keys], maybe remove in future version.
func Keys[K comparable, V any](m map[K]V) []K {
	results := make([]K, 0, len(m))
	for key := range m {
		results = append(results, key)
	}
	return results
}

// Values returns values of map as a slice.
//
// 以切片形式返回 map 的所有值.
//
// Deprecated: As of Go 1.23, please use [maps.Values], maybe remove in future version.
func Values[K comparable, V any](m map[K]V) []V {
	results := make([]V, 0, len(m))
	for _, value := range m {
		results = append(results, value)
	}
	return results
}
