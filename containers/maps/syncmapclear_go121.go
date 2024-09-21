//go:build !go1.23

package maps

import "sync"

func (m *SyncMap[K, V]) Clear() {
	m.Map = sync.Map{}
}
