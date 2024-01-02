package NumArray

import "testing"

func TestSolution(t *testing.T) {
	var want int

	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})

	want = 1
	t.Run("", func(t *testing.T) {
		if got := obj.SumRange(0, 2); got != want {
			t.Errorf(" obj.SumRange() = %v, want %v", got, want)
		}
	})

	want = -1
	t.Run("", func(t *testing.T) {
		if got := obj.SumRange(2, 5); got != want {
			t.Errorf(" obj.SumRange() = %v, want %v", got, want)
		}
	})

	want = -3
	t.Run("", func(t *testing.T) {
		if got := obj.SumRange(0, 5); got != want {
			t.Errorf(" obj.SumRange() = %v, want %v", got, want)
		}
	})
}
