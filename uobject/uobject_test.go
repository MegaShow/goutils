package uobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	assert.Equal(t, 12, Default(12, 100))
	assert.Equal(t, 100, Default(0, 100))
	assert.Equal(t, 12.0, Default(12.0, 100))
	assert.Equal(t, 100.0, Default(0.0, 100))
	assert.Equal(t, "hello", Default("hello", "default"))
	assert.Equal(t, "default", Default("", "default"))
}
