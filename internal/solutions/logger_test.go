package solutions

import (
	"testing"
)

func TestLoggerSmoke1(t *testing.T) {
	obj := LoggerConstructor()
	want := false

	want = true
	t.Run("", func(t *testing.T) {
		if got := obj.ShouldPrintMessage(1, "foo"); got != want {
			t.Errorf(" obj.ShouldPrintMessage() = %v, want = %v", got, want)
		}
	})

	want = true
	t.Run("", func(t *testing.T) {
		if got := obj.ShouldPrintMessage(2, "bar"); got != want {
			t.Errorf(" obj.ShouldPrintMessage() = %v, want = %v", got, want)
		}
	})

	want = false
	t.Run("", func(t *testing.T) {
		if got := obj.ShouldPrintMessage(3, "foo"); got != want {
			t.Errorf(" obj.ShouldPrintMessage() = %v, want = %v", got, want)
		}
	})

	want = false
	t.Run("", func(t *testing.T) {
		if got := obj.ShouldPrintMessage(8, "bar"); got != want {
			t.Errorf(" obj.ShouldPrintMessage() = %v, want = %v", got, want)
		}
	})

	want = false
	t.Run("", func(t *testing.T) {
		if got := obj.ShouldPrintMessage(10, "foo"); got != want {
			t.Errorf(" obj.ShouldPrintMessage() = %v, want = %v", got, want)
		}
	})

	want = true
	t.Run("", func(t *testing.T) {
		if got := obj.ShouldPrintMessage(11, "foo"); got != want {
			t.Errorf(" obj.ShouldPrintMessage() = %v, want = %v", got, want)
		}
	})
}
