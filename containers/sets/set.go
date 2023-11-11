// Package sets implements sets of different type, such as [HashSet].
//
// 包 sets 实现了多种集合数据结构，比如 [HashSet].
package sets

import "go.icytown.com/utils/containers/internal"

// Set is the interface for sets.
//
// 集合接口.
type Set[T comparable] interface {
	Add(item T)
	Contains(item T) bool
	Remove(item T) bool

	internal.Container[T]
}
