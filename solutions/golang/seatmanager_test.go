package golang

import (
	"testing"
)

func TestSeatManagerSmoke1(t *testing.T) {
	t.Parallel()

	obj := SeatManagerConstructor(5)

	want := 1
	if got := obj.Reserve(); got != want {
		t.Errorf("obj.Reserve() = %v, want = %v", got, want)
	}

	want = 2
	if got := obj.Reserve(); got != want {
		t.Errorf("obj.Reserve() = %v, want = %v", got, want)
	}

	obj.Unreserve(2)

	want = 2
	if got := obj.Reserve(); got != want {
		t.Errorf("obj.Reserve() = %v, want = %v", got, want)
	}

	want = 3
	if got := obj.Reserve(); got != want {
		t.Errorf("obj.Reserve() = %v, want = %v", got, want)
	}

	want = 4
	if got := obj.Reserve(); got != want {
		t.Errorf("obj.Reserve() = %v, want = %v", got, want)
	}

	want = 5
	if got := obj.Reserve(); got != want {
		t.Errorf("obj.Reserve() = %v, want = %v", got, want)
	}

	obj.Unreserve(5)
}
