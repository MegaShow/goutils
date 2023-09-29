package assert

import (
	"encoding/json"
	"reflect"
	"testing"
)

// Fail reports a failure.
//
// 报告错误.
func Fail(t *testing.T, msg string, args ...any) bool {
	t.Helper()
	t.Errorf(msg, args...)
	return false
}

// True asserts that value is true.
//
// 断言给定值为 true.
func True(t *testing.T, value bool) bool {
	if !value {
		t.Helper()
		return Fail(t, "Should be true")
	}
	return true
}

// False asserts that value is false.
//
// 断言给定值为 false.
func False(t *testing.T, value bool) bool {
	if value {
		t.Helper()
		return Fail(t, "Should be false")
	}
	return true
}

// Nil asserts that value is nil.
//
// 断言给定值为 nil.
func Nil[T comparable](t *testing.T, value T) bool {
	var zero T
	if value != zero {
		t.Helper()
		actualBytes, _ := json.Marshal(value)
		return Fail(t, "Should be nil, actual: %s", string(actualBytes))
	}
	return true
}

// Zero asserts that value is zero.
//
// 断言给定值为 zero.
func Zero[T comparable](t *testing.T, value T) bool {
	var zero T
	if value != zero {
		t.Helper()
		actualBytes, _ := json.Marshal(value)
		return Fail(t, "Should be zero, actual: %s", string(actualBytes))
	}
	return true
}

// Equal asserts that two objects are equal.
//
// 断言给定两个对象完全相等.
func Equal[T any](t *testing.T, want, actual T) bool {
	if !reflect.DeepEqual(want, actual) {
		t.Helper()
		wantBytes, _ := json.Marshal(want)
		actualBytes, _ := json.Marshal(actual)
		return Fail(t, "Not equal, want: %s, actual: %s", string(wantBytes), string(actualBytes))
	}
	return true
}

// SliceUnorder asserts that two slices are equal without order.
//
// 断言给定两个切片完全相等, 可忽视元素顺序.
func EqualSliceUnorder[T comparable](t *testing.T, want, actual []T) bool {
	// check length
	if len(want) != len(actual) {
		t.Helper()
		wantBytes, _ := json.Marshal(want)
		actualBytes, _ := json.Marshal(actual)
		return Fail(t, "Not equal without order, want: %s, actutal: %s", string(wantBytes), string(actualBytes))
	}

	// check items
	uses := make(map[T]struct{}, len(actual))
	for _, wantItem := range want {
		for _, actualItem := range actual {
			_, ok := uses[actualItem]
			if !ok && reflect.DeepEqual(wantItem, actualItem) {
				uses[actualItem] = struct{}{}
				break
			}
		}
	}
	if len(uses) != len(actual) {
		t.Helper()
		wantBytes, _ := json.Marshal(want)
		actualBytes, _ := json.Marshal(actual)
		return Fail(t, "Not equal without order, want: %s, actutal: %s", string(wantBytes), string(actualBytes))
	}
	return true
}
