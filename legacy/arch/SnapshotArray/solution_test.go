package SnapshotArray

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var obj SnapshotArray
	var want int

	// ---- case 1
	obj = Constructor(3)
	obj.Set(0, 5)
	obj.Snap()
	obj.Set(0, 6)

	want = 5
	t.Run("", func(t *testing.T) {
		if got := obj.Get(0, 0); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})

	// ---- case 2
	obj = Constructor(4)
	obj.Snap()
	obj.Snap()
	obj.Get(3, 1)
	obj.Set(2, 4)
	obj.Snap()
	obj.Set(1, 4)

	want = 0
	t.Run("", func(t *testing.T) {
		if got := obj.Get(0, 0); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})

	// ---- case 3
	obj = Constructor(1)
	obj.Set(0, 15)
	obj.Snap()
	obj.Snap()
	obj.Snap()

	want = 15
	t.Run("", func(t *testing.T) {
		if got := obj.Get(0, 2); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})
	obj.Snap()
	obj.Snap()

	want = 15
	t.Run("", func(t *testing.T) {
		if got := obj.Get(0, 0); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})

	// ---- case 4
	obj = Constructor(3)
	obj.Set(1, 6)
	obj.Snap()
	obj.Snap()
	obj.Set(1, 19)
	obj.Set(0, 4)

	want = 0
	t.Run("", func(t *testing.T) {
		if got := obj.Get(2, 1); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})

	want = 0
	t.Run("", func(t *testing.T) {
		if got := obj.Get(2, 0); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})

	want = 0
	t.Run("", func(t *testing.T) {
		if got := obj.Get(0, 1); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})

	// ---- case 5
	obj = Constructor(1)
	obj.Snap()
	obj.Get(0, 0)
	obj.Snap()
	obj.Get(0, 0)
	obj.Set(0, 13)
	obj.Set(0, 4)
	obj.Set(0, 17)
	obj.Get(0, 0)
	obj.Set(0, 2)
	obj.Get(0, 1)
	obj.Snap()

	want = 2
	t.Run("", func(t *testing.T) {
		if got := obj.Get(0, 2); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})

	obj.Get(0, 0)
	obj.Snap()

	want = 0
	t.Run("", func(t *testing.T) {
		if got := obj.Get(0, 1); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})

	// ---- case 6
	obj = Constructor(3)
	obj.Set(1, 6)
	obj.Snap()
	obj.Snap()
	obj.Set(1, 19)
	obj.Set(0, 4)

	want = 0
	t.Run("", func(t *testing.T) {
		if got := obj.Get(2, 1); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})

	want = 0
	t.Run("", func(t *testing.T) {
		if got := obj.Get(2, 0); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})

	want = 0
	t.Run("", func(t *testing.T) {
		if got := obj.Get(0, 1); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})

	// ---- case 7
	obj = Constructor(1)
	obj.Snap()
	obj.Snap()
	obj.Set(0, 4)
	obj.Snap()
	obj.Get(0, 1)
	obj.Set(0, 12)
	obj.Get(0, 1)
	obj.Snap()

	want = 12
	t.Run("", func(t *testing.T) {
		if got := obj.Get(0, 3); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})
}
