package assert

import (
	"fmt"
	"strings"
	"testing"
)

type recordingTesting struct {
	messages []string
	helpers  int
}

func (r *recordingTesting) Errorf(format string, args ...any) {
	r.messages = append(r.messages, fmt.Sprintf(format, args...))
}

func (r *recordingTesting) Helper() { r.helpers++ }

func TestElementsMatch(t *testing.T) {
	t.Parallel()
	r := &recordingTesting{}
	if !ElementsMatch(r, []int{1, 1, 2}, []int{2, 1, 1}) || len(r.messages) != 0 {
		t.Fatal("successful assertion recorded an error")
	}
	if ElementsMatch(r, []int{1, 1}, []int{1, 2}) || len(r.messages) != 1 || r.helpers != 1 || r.messages[0] == "" {
		t.Fatal("failed assertion did not record an error")
	}

	r = &recordingTesting{}
	if ElementsMatch(r, []int{1}, []int{}) || len(r.messages) != 1 {
		t.Fatal("length mismatch did not record an error")
	}
}

func TestEqual(t *testing.T) {
	t.Parallel()
	r := &recordingTesting{}
	if !Equal(r, 1, 1) || len(r.messages) != 0 {
		t.Fatal("successful assertion recorded an error")
	}
	if Equal(r, 1, 2) || len(r.messages) != 1 || r.helpers != 1 || r.messages[0] == "" {
		t.Fatal("failed assertion did not record an error")
	}
}

func TestFalse(t *testing.T) {
	t.Parallel()
	r := &recordingTesting{}
	if !False(r, false) || len(r.messages) != 0 {
		t.Fatal("successful assertion recorded an error")
	}
	if False(r, true) || len(r.messages) != 1 || r.helpers != 1 || r.messages[0] == "" {
		t.Fatal("failed assertion did not record an error")
	}
}

func TestNil(t *testing.T) {
	t.Parallel()
	r := &recordingTesting{}
	if !Nil(r, (*int)(nil)) || len(r.messages) != 0 {
		t.Fatal("successful assertion recorded an error")
	}
	if Nil(r, 1) || len(r.messages) != 1 || r.helpers != 1 || r.messages[0] == "" {
		t.Fatal("failed assertion did not record an error")
	}
}

func TestNotEqual(t *testing.T) {
	t.Parallel()
	r := &recordingTesting{}
	if !NotEqual(r, 1, 2) || len(r.messages) != 0 {
		t.Fatal("successful assertion recorded an error")
	}
	if NotEqual(r, 1, 1) || len(r.messages) != 1 || r.helpers != 1 || r.messages[0] == "" {
		t.Fatal("failed assertion did not record an error")
	}
}

func TestNotNil(t *testing.T) {
	t.Parallel()
	r := &recordingTesting{}
	if !NotNil(r, 1) || len(r.messages) != 0 {
		t.Fatal("successful assertion recorded an error")
	}
	if NotNil(r, (*int)(nil)) || len(r.messages) != 1 || r.helpers != 1 || r.messages[0] == "" {
		t.Fatal("failed assertion did not record an error")
	}
}

func TestNotPanics(t *testing.T) {
	t.Parallel()
	r := &recordingTesting{}
	if !NotPanics(r, func() {}) || len(r.messages) != 0 {
		t.Fatal("successful assertion recorded an error")
	}
	if NotPanics(r, func() { panic("test panic") }) || len(r.messages) != 1 || r.helpers != 1 || r.messages[0] == "" {
		t.Fatal("failed assertion did not record an error")
	}
}

func TestPanics(t *testing.T) {
	t.Parallel()
	r := &recordingTesting{}
	if !Panics(r, func() { panic("test panic") }) || len(r.messages) != 0 {
		t.Fatal("successful assertion recorded an error")
	}
	if Panics(r, func() {}) || len(r.messages) != 1 || r.helpers != 1 || r.messages[0] == "" {
		t.Fatal("failed assertion did not record an error")
	}
}

func TestTrue(t *testing.T) {
	t.Parallel()
	r := &recordingTesting{}
	if !True(r, true) || len(r.messages) != 0 {
		t.Fatal("successful assertion recorded an error")
	}
	if True(r, false) || len(r.messages) != 1 || r.helpers != 1 || r.messages[0] == "" {
		t.Fatal("failed assertion did not record an error")
	}
}

func TestZero(t *testing.T) {
	t.Parallel()
	r := &recordingTesting{}
	if !Zero(r, 0) || len(r.messages) != 0 {
		t.Fatal("successful assertion recorded an error")
	}
	if Zero(r, 1) || len(r.messages) != 1 || r.helpers != 1 || r.messages[0] == "" {
		t.Fatal("failed assertion did not record an error")
	}
}

func Test_isNil(t *testing.T) {
	t.Parallel()
	for _, value := range []any{nil, (*int)(nil), []int(nil), map[int]int(nil), (chan int)(nil), (func())(nil)} {
		if !isNil(value) {
			t.Fatalf("isNil(%T) = false", value)
		}
	}
	for _, value := range []any{0, "", []int{}, map[int]int{}, new(int)} {
		if isNil(value) {
			t.Fatalf("isNil(%T) = true", value)
		}
	}
}

func Test_runAndRecover(t *testing.T) {
	t.Parallel()
	called := false
	value, stack, recovered := runAndRecover(func() { called = true })
	True(t, called)
	Nil(t, value)
	Equal(t, "", stack)
	False(t, recovered)

	value, stack, recovered = runAndRecover(func() { panic("test panic") })
	Equal(t, any("test panic"), value)
	True(t, strings.Contains(stack, "runAndRecover"))
	True(t, recovered)
}
