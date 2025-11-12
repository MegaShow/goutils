// Package assert provides assertion utilities for testing.
//
// 包 assert 提供了测试用的断言工具.
package assert

import (
	"reflect"
	"runtime/debug"
)

type Testing interface {
	Errorf(format string, args ...any)
	Helper()
}

func ElementsMatch[T any](t Testing, expected, actual []T) bool {
	if len(expected) != len(actual) {
		t.Helper()
		t.Errorf("Should have same length:\n"+
			"Expected: %d\n"+
			"Actual:   %d", len(expected), len(actual))
		return false
	}

	matched := make([]bool, len(actual))
	for _, expectedValue := range expected {
		found := false
		for idx, actualValue := range actual {
			if !matched[idx] && reflect.DeepEqual(expectedValue, actualValue) {
				matched[idx] = true
				found = true
				break
			}
		}
		if !found {
			t.Helper()
			t.Errorf("Element not found in actual slice:\n"+
				"Missing element: %#v\n"+
				"Expected: %#v\n"+
				"Actual:   %#v", expectedValue, expected, actual)
			return false
		}
	}
	return true
}

func Equal[T any](t Testing, expected, actual T) bool {
	if !reflect.DeepEqual(expected, actual) {
		t.Helper()
		t.Errorf("Should be equal:\n"+
			"Expected: %#v\n"+
			"Actual:   %#v", expected, actual)
		return false
	}
	return true
}

func False(t Testing, condition bool) bool {
	if condition {
		t.Helper()
		t.Errorf("Should be false")
		return false
	}
	return true
}

func Nil(t Testing, value any) bool {
	if !isNil(value) {
		t.Helper()
		t.Errorf("Should be nil, but got %#v", value)
		return false
	}
	return true
}

func NotEqual[T any](t Testing, expected, actual T) bool {
	if reflect.DeepEqual(expected, actual) {
		t.Helper()
		t.Errorf("Should be not equal, value: %#v", actual)
		return false
	}
	return true
}

func NotNil(t Testing, value any) bool {
	if isNil(value) {
		t.Helper()
		t.Errorf("Should be not nil, but got nil")
		return false
	}
	return true

}

func NotPanics(t Testing, f func()) bool {
	x, stack, recovered := runAndRecover(f)
	if recovered {
		t.Helper()
		t.Errorf("Should not panic:\n"+
			"Panic value: %v\n"+
			"Panic stack: %s", x, stack)
		return false
	}
	return true
}

func Panics(t Testing, f func()) bool {
	_, _, recovered := runAndRecover(f)
	if !recovered {
		t.Helper()
		t.Errorf("Should panic")
		return false
	}
	return true
}

func True(t Testing, condition bool) bool {
	if !condition {
		t.Helper()
		t.Errorf("Should be true")
		return false
	}
	return true
}

func Zero[T any](t Testing, value T) bool {
	var zero T
	if !reflect.DeepEqual(value, zero) {
		t.Helper()
		t.Errorf("Should be zero value, but got %#v", value)
		return false
	}
	return true
}

func isNil(obj any) bool {
	if obj == nil {
		return true
	}
	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		return value.IsNil()
	default:
		return false
	}
}

func runAndRecover(f func()) (x any, stack string, recovered bool) {
	recovered = true
	defer func() {
		x = recover()
		if recovered {
			stack = string(debug.Stack())
		}
	}()
	f()
	recovered = false
	return
}
