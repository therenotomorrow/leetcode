package MyStack

import "testing"

func TestSolution(t *testing.T) {
	var want int

	// ---- case 1
	obj1 := Constructor()

	obj1.Push(1) // queue is: [1]
	obj1.Push(2) // queue is: [1, 2] (leftmost is front of the queue)

	want = 2
	t.Run("", func(t *testing.T) {
		if got := obj1.Top(); got != want {
			t.Errorf(" obj.Top() = %v, want %v", got, want)
		}
	})

	want = 2
	t.Run("", func(t *testing.T) {
		if got := obj1.Pop(); got != want {
			t.Errorf(" obj.Pop() = %v, want %v", got, want)
		}
	})

	t.Run("", func(t *testing.T) {
		if got := obj1.Empty(); got != false {
			t.Errorf(" obj.Pop() = %v, want %v", got, false)
		}
	})
}
