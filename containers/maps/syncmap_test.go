package maps

import (
	"sync"
	"testing"

	"go.icytown.com/utils/internal/assert"
)

func TestSyncMap_Clear(t *testing.T) {
	t.Parallel()
	var m SyncMap[int, int]
	m.Clear()
	var wg sync.WaitGroup
	for worker := range 4 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range 100 {
				m.Store(worker, i)
				_, _ = m.Load(worker)
				m.Clear()
			}
		}()
	}
	wg.Wait()
	m.Clear()
	count := 0
	m.Range(func(_, _ int) bool { count++; return true })
	assert.Equal(t, 0, count)
	m.Store(1, 2)
	value, ok := m.Load(1)
	assert.True(t, ok)
	assert.Equal(t, 2, value)
}

func TestSyncMap_CompareAndDelete(t *testing.T) {
	t.Parallel()
	var m SyncMap[string, int]
	assert.False(t, m.CompareAndDelete("key", 1))
	m.Store("key", 1)
	assert.False(t, m.CompareAndDelete("key", 2))
	assert.True(t, m.CompareAndDelete("key", 1))
	_, ok := m.Load("key")
	assert.False(t, ok)
}

func TestSyncMap_CompareAndSwap(t *testing.T) {
	t.Parallel()
	var m SyncMap[string, int]
	assert.False(t, m.CompareAndSwap("key", 1, 2))
	m.Store("key", 1)
	assert.False(t, m.CompareAndSwap("key", 0, 2))
	value, _ := m.Load("key")
	assert.Equal(t, 1, value)
	assert.True(t, m.CompareAndSwap("key", 1, 2))
	value, _ = m.Load("key")
	assert.Equal(t, 2, value)
}

func TestSyncMap_Delete(t *testing.T) {
	t.Parallel()
	var m SyncMap[string, int]
	m.Delete("missing")
	m.Store("key", 1)
	m.Delete("key")
	_, ok := m.Load("key")
	assert.False(t, ok)
}

func TestSyncMap_Load(t *testing.T) {
	t.Parallel()
	var m SyncMap[string, any]
	value, ok := m.Load("missing")
	assert.Nil(t, value)
	assert.False(t, ok)
	m.Store("nil", nil)
	value, ok = m.Load("nil")
	assert.Nil(t, value)
	assert.True(t, ok)
	var pointer *int
	m.Store("pointer", pointer)
	value, ok = m.Load("pointer")
	assert.Equal(t, any(pointer), value)
	assert.True(t, ok)
}

func TestSyncMap_LoadAndDelete(t *testing.T) {
	t.Parallel()
	var m SyncMap[string, any]
	value, loaded := m.LoadAndDelete("missing")
	assert.Nil(t, value)
	assert.False(t, loaded)
	m.Store("key", nil)
	value, loaded = m.LoadAndDelete("key")
	assert.Nil(t, value)
	assert.True(t, loaded)
	_, loaded = m.Load("key")
	assert.False(t, loaded)
	m.Store("key", 42)
	value, loaded = m.LoadAndDelete("key")
	assert.Equal(t, any(42), value)
	assert.True(t, loaded)
}

func TestSyncMap_LoadOrStore(t *testing.T) {
	t.Parallel()
	var m SyncMap[string, any]
	value, loaded := m.LoadOrStore("key", nil)
	assert.Nil(t, value)
	assert.False(t, loaded)
	value, loaded = m.LoadOrStore("key", 42)
	assert.Nil(t, value)
	assert.True(t, loaded)
	value, loaded = m.LoadOrStore("other", 42)
	assert.Equal(t, any(42), value)
	assert.False(t, loaded)
}

func TestSyncMap_Range(t *testing.T) {
	t.Parallel()
	var m SyncMap[any, any]
	m.Store(nil, nil)
	m.Store("key", 42)
	values := make(map[any]any)
	m.Range(func(key, value any) bool {
		values[key] = value
		return true
	})
	assert.Equal(t, map[any]any{nil: nil, "key": 42}, values)
	calls := 0
	m.Range(func(_, _ any) bool { calls++; return false })
	assert.Equal(t, 1, calls)
}

func TestSyncMap_Swap(t *testing.T) {
	t.Parallel()
	var m SyncMap[string, any]
	value, loaded := m.Swap("key", nil)
	assert.Nil(t, value)
	assert.False(t, loaded)
	value, loaded = m.Swap("key", 42)
	assert.Nil(t, value)
	assert.True(t, loaded)
	value, loaded = m.Swap("key", 43)
	assert.Equal(t, any(42), value)
	assert.True(t, loaded)
	value, loaded = m.Load("key")
	assert.Equal(t, any(43), value)
	assert.True(t, loaded)
}
