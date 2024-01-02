package LRUCache

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var want int

	obj := Constructor(2)

	obj.Put(1, 1) // cache is {1=1}
	obj.Put(2, 2) // cache is {1=1, 2=2}

	want = 1
	t.Run("", func(t *testing.T) {
		if got := obj.Get(1); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})

	obj.Put(3, 3) // LRU key was 2, evicts key 2, cache is {1=1, 3=3}

	want = -1
	t.Run("", func(t *testing.T) {
		if got := obj.Get(2); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})

	obj.Put(4, 4) // LRU key was 1, evicts key 1, cache is {4=4, 3=3}

	want = -1
	t.Run("", func(t *testing.T) {
		if got := obj.Get(1); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})

	want = 3
	t.Run("", func(t *testing.T) {
		if got := obj.Get(3); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})

	want = 4
	t.Run("", func(t *testing.T) {
		if got := obj.Get(4); got != want {
			t.Errorf(" obj.Get() = %v, want %v", got, want)
		}
	})
}
