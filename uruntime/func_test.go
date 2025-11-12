package uruntime

import (
	"testing"

	"go.icytown.com/utils/internal/assert"
)

var __fnVar func(a, b int) int

func __add(a, b int) int {
	return a + b
}

type __struct struct{}

func (obj __struct) Method() {}
func (obj __struct) method() {}

func TestGetFuncFullName(t *testing.T) {
	assert.Equal(t, "go.icytown.com/utils/uruntime.GetFuncFullName", GetFuncFullName(GetFuncFullName))
	assert.Equal(t, "go.icytown.com/utils/uruntime.GetFuncName", GetFuncFullName(GetFuncName))
	assert.Equal(t, "go.icytown.com/utils/uruntime.TestGetFuncFullName.func1", GetFuncFullName(func() {}))
	assert.Equal(t, "go.icytown.com/utils/uruntime.__add", GetFuncFullName(__add))
	assert.Equal(t, "", GetFuncFullName(__fnVar))
	assert.Equal(t, "", GetFuncFullName(nil))
	assert.Equal(t, "", GetFuncFullName(1))
	assert.Equal(t, "go.icytown.com/utils/uruntime.__struct.Method", GetFuncFullName(__struct.Method))
	assert.Equal(t, "go.icytown.com/utils/uruntime.__struct.method", GetFuncFullName(__struct.method))
}

func TestGetFuncName(t *testing.T) {
	assert.Equal(t, "GetFuncFullName", GetFuncName(GetFuncFullName))
	assert.Equal(t, "GetFuncName", GetFuncName(GetFuncName))
	assert.Equal(t, "func1", GetFuncName(func() {}))
	assert.Equal(t, "__add", GetFuncName(__add))
	assert.Equal(t, "", GetFuncName(__fnVar))
	assert.Equal(t, "", GetFuncName(nil))
	assert.Equal(t, "", GetFuncName(1))
	assert.Equal(t, "Method", GetFuncName(__struct.Method))
	assert.Equal(t, "method", GetFuncName(__struct.method))
}
