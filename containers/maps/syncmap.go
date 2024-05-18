package maps

import "sync"

// SyncMap is a wrapper for sync.Map, which supports generics.
//
// sync.Map 的封装, 支持泛型.
type SyncMap[K comparable, V any] struct {
	sync.Map
}

// CompareAndDelete deletes the entry for key if its value is equal to old.
//
// 删除对应键值如果值和给定的 old 变量相等.
func (m *SyncMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return m.Map.CompareAndDelete(key, old)
}

// CompareAndSwap swaps the old and new values for key if the value stored in the map is equal to old.
//
// 交换指定键的新旧值如果存储的旧值和给定的 old 变量相等.
func (m *SyncMap[K, V]) CompareAndSwap(key K, old, new V) bool {
	return m.Map.CompareAndSwap(key, old, new)
}

// Delete deletes the value for a key.
//
// 删除指定键.
func (m *SyncMap[K, V]) Delete(key K) {
	m.Map.Delete(key)
}

// Load returns the value stored in the map for a key.
//
// 返回指定键的值.
func (m *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.Map.Load(key)
	if !ok {
		return
	}
	value, ok = v.(V)
	return
}

// LoadAndDelete deletes the value for a key, returning the previous value.
//
// 删除指定键, 并返回该键的值.
func (m *SyncMap[K, V]) LoadAndDelete(key K) (actual V, loaded bool) {
	v, loaded := m.Map.LoadAndDelete(key)
	return v.(V), loaded
}

// LoadOrStore returns the existing value for the key if present. Otherwise, it stores and returns the given value.
//
// 如果指定键存在则返回对应的值, 否则则存储指定值并返回它.
func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	v, loaded := m.Map.LoadOrStore(key, value)
	return v.(V), loaded
}

// Range calls f sequentially for each key and value.
//
// 遍历调用 f 函数, 并传入每一对键值.
func (m *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	m.Map.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

// Store sets the value for a key.
//
// 存储指定键值.
func (m *SyncMap[K, V]) Store(key K, value V) {
	m.Map.Store(key, value)
}

// Swap swaps the value for a key and returns the previous value.
//
// 存储指定键的新值, 并返回旧值.
func (m *SyncMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	v, loaded := m.Map.Swap(key, value)
	return v.(V), loaded
}
