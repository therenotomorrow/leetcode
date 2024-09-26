package golang

import "testing"

func TestMyCalendarSmoke1(t *testing.T) {
	t.Parallel()

	obj := MyCalendarConstructor()

	want := true
	if got := obj.Book(10, 20); got != want {
		t.Errorf("obj.Book() = %v, want = %v", got, want)
	}

	want = false
	if got := obj.Book(15, 25); got != want {
		t.Errorf("obj.Book() = %v, want = %v", got, want)
	}

	want = true
	if got := obj.Book(20, 30); got != want {
		t.Errorf("obj.Book() = %v, want = %v", got, want)
	}
}
