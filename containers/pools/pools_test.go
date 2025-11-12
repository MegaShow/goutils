package pools

import (
	"testing"

	"go.icytown.com/utils/internal/assert"
)

func TestNewPool(t *testing.T) {
	type s struct{}
	p := NewPool[*s](nil)
	assert.NotNil(t, p)
	p = NewPool(func() *s { return &s{} })
	assert.NotNil(t, p)
}

func TestPool_Get(t *testing.T) {
	type s struct{ v int }
	p := NewPool[*s](nil)
	assert.Nil(t, p.Get())

	var count int
	p = NewPool(func() *s {
		count++
		return &s{v: count}
	})
	assert.Equal(t, 1, p.Get().v)
	assert.Equal(t, 2, p.Get().v)
	assert.Equal(t, 3, p.Get().v)
}

func TestPool_Put(t *testing.T) {
	type s struct{ v int }
	p := NewPool(func() *s { return &s{} })
	for i := 0; i < 100; i++ {
		p.Put(&s{v: 1})
	}
	assert.Equal(t, 1, p.Get().v)
}
