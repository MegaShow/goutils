package umap

// Keys returns keys of map as a slice.
//
// 以切片形式返回 map 的所有键.
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
func Values[K comparable, V any](m map[K]V) []V {
	results := make([]V, 0, len(m))
	for _, value := range m {
		results = append(results, value)
	}
	return results
}
