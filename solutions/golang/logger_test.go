package golang

import (
	"testing"
)

func TestLoggerSmoke1(t *testing.T) {
	t.Parallel()

	obj := LoggerConstructor()

	want := true
	if got := obj.ShouldPrintMessage(1, "foo"); got != want {
		t.Errorf(" obj.ShouldPrintMessage() = %v, want = %v", got, want)
	}

	want = true
	if got := obj.ShouldPrintMessage(2, "bar"); got != want {
		t.Errorf(" obj.ShouldPrintMessage() = %v, want = %v", got, want)
	}

	want = false
	if got := obj.ShouldPrintMessage(3, "foo"); got != want {
		t.Errorf(" obj.ShouldPrintMessage() = %v, want = %v", got, want)
	}

	want = false
	if got := obj.ShouldPrintMessage(8, "bar"); got != want {
		t.Errorf(" obj.ShouldPrintMessage() = %v, want = %v", got, want)
	}

	want = false
	if got := obj.ShouldPrintMessage(10, "foo"); got != want {
		t.Errorf(" obj.ShouldPrintMessage() = %v, want = %v", got, want)
	}

	want = true
	if got := obj.ShouldPrintMessage(11, "foo"); got != want {
		t.Errorf(" obj.ShouldPrintMessage() = %v, want = %v", got, want)
	}
}
