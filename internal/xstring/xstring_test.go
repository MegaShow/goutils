package xstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesToString(t *testing.T) {
	assert.Equal(t, "", BytesToString([]byte{}))
	assert.Equal(t, "123456", BytesToString([]byte{'1', '2', '3', '4', '5', '6'}))
}

func TestStringToBytes(t *testing.T) {
	assert.Equal(t, []byte(nil), StringToBytes(""))
	assert.Equal(t, []byte{'1', '2', '3', '4', '5', '6'}, StringToBytes("123456"))
}
