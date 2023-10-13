package RecentCounter

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var want int

	obj := Constructor()

	want = 1
	t.Run("", func(t *testing.T) {
		if got := obj.Ping(1); got != want {
			t.Errorf(" obj.Ping() = %v, want %v", got, want)
		}
	})

	want = 2
	t.Run("", func(t *testing.T) {
		if got := obj.Ping(100); got != want {
			t.Errorf(" obj.Ping() = %v, want %v", got, want)
		}
	})

	want = 3
	t.Run("", func(t *testing.T) {
		if got := obj.Ping(3001); got != want {
			t.Errorf(" obj.Ping() = %v, want %v", got, want)
		}
	})

	want = 3
	t.Run("", func(t *testing.T) {
		if got := obj.Ping(3002); got != want {
			t.Errorf(" obj.Ping() = %v, want %v", got, want)
		}
	})
}
