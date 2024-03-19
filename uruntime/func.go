// Package uruntime provides utils of runtime.
//
// 包 uruntime 提供了运行时相关的工具.
package uruntime

import (
	"reflect"
	"runtime"
	"strings"
)

// GetFuncFullName returns full name of Golang function or method with package path,
// e.g. go.icytown.com/utils/runtimeutils.GetFuncFullName.
//
// 获取 Golang 的函数或方法包括包路径的名称, 比如 go.icytown.com/utils/runtimeutils.GetFuncFullName.
func GetFuncFullName(fn any) string {
	fnValue := reflect.ValueOf(fn)
	if fnValue.Kind() != reflect.Func || fnValue.IsNil() {
		return ""
	}
	return runtime.FuncForPC(fnValue.Pointer()).Name()
}

// GetFuncName returns name of Golang function or method, e.g. GetFuncName.
//
// 获取 Golang 的函数或方法名, 比如 GetFuncName.
func GetFuncName(fn any) string {
	fnName := GetFuncFullName(fn)
	idx := strings.LastIndexByte(fnName, '.')
	if idx == -1 {
		return fnName
	}
	return fnName[idx+1:]
}
