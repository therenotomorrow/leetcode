package golang

import "testing"

func TestMyQueueSmoke1(t *testing.T) {
	t.Parallel()

	obj := MyQueueConstructor()

	obj.Push(1) // queue is: [1]
	obj.Push(2) // queue is: [1, 2] (leftmost is front of the queue)

	want := 1
	if got := obj.Peek(); got != want {
		t.Errorf("obj.Peek() = %v, want = %v", got, want)
	}

	want = 1
	if got := obj.Pop(); got != want {
		t.Errorf("obj.Pop() = %v, want = %v", got, want)
	}

	if got := obj.Empty(); got {
		t.Errorf("obj.Empty() = %v, want = %v", got, false)
	}
}
