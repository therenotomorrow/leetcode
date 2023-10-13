package MovingAverage

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var want float64
	obj := Constructor(3)

	want = 1.0
	t.Run("", func(t *testing.T) {
		if got := obj.Next(1); got != want {
			t.Errorf(" obj.Next() = %v, want %v", got, want)
		}
	})

	want = 5.5
	t.Run("", func(t *testing.T) {
		if got := obj.Next(10); got != want {
			t.Errorf(" obj.Next() = %v, want %v", got, want)
		}
	})

	obj.Next(3)
	want = 6.0
	t.Run("", func(t *testing.T) {
		if got := obj.Next(5); got != want {
			t.Errorf(" obj.Next() = %v, want %v", got, want)
		}
	})
}
