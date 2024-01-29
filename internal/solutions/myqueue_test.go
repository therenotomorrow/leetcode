package solutions

import "testing"

func TestMyQueueSmoke1(t *testing.T) {
	var want int

	obj := MyQueueConstructor()

	obj.Push(1) // queue is: [1]
	obj.Push(2) // queue is: [1, 2] (leftmost is front of the queue)

	want = 1

	t.Run("", func(t *testing.T) {
		if got := obj.Peek(); got != want {
			t.Errorf(" obj.Peek() = %v, want = %v", got, want)
		}
	})

	want = 1

	t.Run("", func(t *testing.T) {
		if got := obj.Pop(); got != want {
			t.Errorf(" obj.Pop() = %v, want = %v", got, want)
		}
	})

	t.Run("", func(t *testing.T) {
		if got := obj.Empty(); got {
			t.Errorf(" obj.Empty() = %v, want = %v", got, false)
		}
	})
}
