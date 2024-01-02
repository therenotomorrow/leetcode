package solutions

import (
	"testing"
)

func TestSeatManagerSmoke1(t *testing.T) {
	obj := SeatManagerConstructor(5)
	want := 0

	want = 1

	t.Run("", func(t *testing.T) {
		if got := obj.Reserve(); got != want {
			t.Errorf(" obj.Reserve() = %v, want = %v", got, want)
		}
	})

	want = 2

	t.Run("", func(t *testing.T) {
		if got := obj.Reserve(); got != want {
			t.Errorf(" obj.Reserve() = %v, want = %v", got, want)
		}
	})

	obj.Unreserve(2)

	want = 2

	t.Run("", func(t *testing.T) {
		if got := obj.Reserve(); got != want {
			t.Errorf(" obj.Reserve() = %v, want = %v", got, want)
		}
	})

	want = 3

	t.Run("", func(t *testing.T) {
		if got := obj.Reserve(); got != want {
			t.Errorf(" obj.Reserve() = %v, want = %v", got, want)
		}
	})

	want = 4

	t.Run("", func(t *testing.T) {
		if got := obj.Reserve(); got != want {
			t.Errorf(" obj.Reserve() = %v, want = %v", got, want)
		}
	})

	want = 5

	t.Run("", func(t *testing.T) {
		if got := obj.Reserve(); got != want {
			t.Errorf(" obj.Reserve() = %v, want = %v", got, want)
		}
	})

	obj.Unreserve(5)
}
