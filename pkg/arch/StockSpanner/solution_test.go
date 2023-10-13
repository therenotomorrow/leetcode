package StockSpanner

import "testing"

func TestSolution(t *testing.T) {
	var want int

	// ---- case 1
	obj := Constructor()

	want = 1
	t.Run("", func(t *testing.T) {
		if got := obj.Next(100); got != want {
			t.Errorf(" obj.Next() = %v, want %v", got, want)
		}
	})

	want = 1
	t.Run("", func(t *testing.T) {
		if got := obj.Next(80); got != want {
			t.Errorf(" obj.Next() = %v, want %v", got, want)
		}
	})

	want = 1
	t.Run("", func(t *testing.T) {
		if got := obj.Next(60); got != want {
			t.Errorf(" obj.Next() = %v, want %v", got, want)
		}
	})

	want = 2
	t.Run("", func(t *testing.T) {
		if got := obj.Next(70); got != want {
			t.Errorf(" obj.Next() = %v, want %v", got, want)
		}
	})

	want = 1
	t.Run("", func(t *testing.T) {
		if got := obj.Next(60); got != want {
			t.Errorf(" obj.Next() = %v, want %v", got, want)
		}
	})

	want = 4
	t.Run("", func(t *testing.T) {
		if got := obj.Next(75); got != want {
			t.Errorf(" obj.Next() = %v, want %v", got, want)
		}
	})

	want = 6
	t.Run("", func(t *testing.T) {
		if got := obj.Next(85); got != want {
			t.Errorf(" obj.Next() = %v, want %v", got, want)
		}
	})
}
