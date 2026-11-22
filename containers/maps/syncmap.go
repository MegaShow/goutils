package maps

import "sync"

// SyncMap is a wrapper for sync.Map, which supports generics.
//
// sync.Map 的封装, 支持泛型.
type SyncMap[K comparable, V any] struct {
	items sync.Map
}

// Clear deletes all entries from the map.
//
// 清空所有键值
func (m *SyncMap[K, V]) Clear() {
	m.items.Clear()
}

// CompareAndDelete deletes the entry for key if its value is equal to oldValue.
//
// 删除对应键值如果值和给定的 oldValue 变量相等
func (m *SyncMap[K, V]) CompareAndDelete(key K, oldValue V) (deleted bool) {
	return m.items.CompareAndDelete(key, oldValue)
}

// CompareAndSwap swaps the old and new values for key if the value stored in the map is equal to oldValue.
//
// 交换指定键的新旧值如果存储的旧值和给定的 oldValue 变量相等
func (m *SyncMap[K, V]) CompareAndSwap(key K, oldValue, newValue V) bool {
	return m.items.CompareAndSwap(key, oldValue, newValue)
}

// Delete deletes the value for a key.
//
// 删除指定键.
func (m *SyncMap[K, V]) Delete(key K) {
	m.items.Delete(key)
}

// Load returns the value stored in the map for a key.
//
// 返回指定键的值.
func (m *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.items.Load(key)
	value, _ = v.(V)
	return value, ok
}

// LoadAndDelete deletes the value for a key, returning the previous value.
//
// 删除指定键, 并返回该键的值.
func (m *SyncMap[K, V]) LoadAndDelete(key K) (actual V, loaded bool) {
	v, loaded := m.items.LoadAndDelete(key)
	actual, _ = v.(V)
	return actual, loaded
}

// LoadOrStore returns the existing value for the key if present. Otherwise, it stores and returns the given value.
//
// 如果指定键存在则返回对应的值, 否则则存储指定值并返回它.
func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	v, loaded := m.items.LoadOrStore(key, value)
	actual, _ = v.(V)
	return actual, loaded
}

// Range calls f sequentially for each key and value.
//
// 遍历调用 f 函数, 并传入每一对键值.
func (m *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	m.items.Range(func(key, value any) bool {
		k, _ := key.(K)
		v, _ := value.(V)
		return f(k, v)
	})
}

// Store sets the value for a key.
//
// 存储指定键值.
func (m *SyncMap[K, V]) Store(key K, value V) {
	m.items.Store(key, value)
}

// Swap swaps the value for a key and returns the previous value.
//
// 存储指定键的新值, 并返回旧值.
func (m *SyncMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	v, loaded := m.items.Swap(key, value)
	previous, _ = v.(V)
	return previous, loaded
}
