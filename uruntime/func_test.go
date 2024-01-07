package uruntime

import "testing"

var __fnVar func(a, b int) int

func __add(a, b int) int {
	return a + b
}

type __struct struct{}

func (obj __struct) Method() {}
func (obj __struct) method() {}

func TestGetFuncFullName(t *testing.T) {
	tests := []struct {
		fn   any
		want string
	}{
		{fn: GetFuncFullName, want: "go.icytown.com/utils/runtimeutils.GetFuncFullName"},
		{fn: GetFuncName, want: "go.icytown.com/utils/runtimeutils.GetFuncName"},
		{fn: func() {}, want: "go.icytown.com/utils/runtimeutils.TestGetFuncFullName.func1"},
		{fn: __add, want: "go.icytown.com/utils/runtimeutils.__add"},
		{fn: __fnVar, want: ""},
		{fn: nil, want: ""},
		{fn: 1, want: ""},
		{fn: __struct.Method, want: "go.icytown.com/utils/runtimeutils.__struct.Method"},
		{fn: __struct.method, want: "go.icytown.com/utils/runtimeutils.__struct.method"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := GetFuncFullName(tt.fn); got != tt.want {
				t.Errorf("GetFuncFullName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFuncName(t *testing.T) {
	tests := []struct {
		fn   any
		want string
	}{
		{fn: GetFuncFullName, want: "GetFuncFullName"},
		{fn: GetFuncName, want: "GetFuncName"},
		{fn: func() {}, want: "func1"},
		{fn: __add, want: "__add"},
		{fn: __fnVar, want: ""},
		{fn: nil, want: ""},
		{fn: 1, want: ""},
		{fn: __struct.Method, want: "Method"},
		{fn: __struct.method, want: "method"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := GetFuncName(tt.fn); got != tt.want {
				t.Errorf("GetFuncName() = %v, want %v", got, tt.want)
			}
		})
	}
}
