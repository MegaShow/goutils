//go:build go1.23

package maps

func (m *SyncMap[K, V]) Clear() {
	m.Map.Clear()
}
