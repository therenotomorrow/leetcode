package MyQueue

import "testing"

func TestSolution(t *testing.T) {
	var want int

	// ---- case 1
	obj1 := Constructor()

	obj1.Push(1) // queue is: [1]
	obj1.Push(2) // queue is: [1, 2] (leftmost is front of the queue)

	want = 1
	t.Run("", func(t *testing.T) {
		if got := obj1.Peek(); got != want {
			t.Errorf(" obj.Peek() = %v, want %v", got, want)
		}
	})

	want = 1
	t.Run("", func(t *testing.T) {
		if got := obj1.Pop(); got != want {
			t.Errorf(" obj.Pop() = %v, want %v", got, want)
		}
	})

	// ---- case 2
}
