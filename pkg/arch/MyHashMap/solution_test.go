package MyHashMap

import "testing"

func TestSolution(t *testing.T) {
	var want int

	obj := Constructor()

	obj.Put(1, 1)
	obj.Put(2, 2)

	want = 1
	t.Run("", func(t *testing.T) {
		if got := obj.Get(1); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})

	want = -1
	t.Run("", func(t *testing.T) {
		if got := obj.Get(3); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})

	obj.Put(2, 1)

	want = 1
	t.Run("", func(t *testing.T) {
		if got := obj.Get(2); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})

	obj.Remove(2)

	want = -1
	t.Run("", func(t *testing.T) {
		if got := obj.Get(2); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})
}
