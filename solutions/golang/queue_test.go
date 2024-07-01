package golang_test

import (
	"testing"

	"github.com/therenotomorrow/leetcode/solutions/golang"
)

func TestQueueEnqueueDequeue(t *testing.T) {
	t.Parallel()

	que := golang.NewQueue[int]()

	que.Enqueue(1)
	que.Enqueue(2)
	que.Enqueue(3)

	if got, ok := que.Dequeue(); got != 1 || !ok {
		t.Errorf("Dequeue() = (%v, %v), want = (%v, %v)", got, ok, 1, true)
	}

	if got, ok := que.Dequeue(); got != 2 || !ok {
		t.Errorf("Dequeue() = (%v, %v), want = (%v, %v)", got, ok, 2, true)
	}

	if got, ok := que.Dequeue(); got != 3 || !ok {
		t.Errorf("Dequeue() = (%v, %v), want = (%v, %v)", got, ok, 3, true)
	}

	if got, ok := que.Dequeue(); got != 0 || ok {
		t.Errorf("Dequeue() = (%v, %v), want = (%v, %v)", got, ok, 0, false)
	}
}

func TestQueuePeek(t *testing.T) {
	t.Parallel()

	que := golang.NewQueue[int]()

	que.Enqueue(42)

	if got, ok := que.Peek(); got != 42 || !ok {
		t.Errorf("Peek() = (%v, %v), want = (%v, %v)", got, ok, 42, true)
	}

	que.Dequeue()

	if got, ok := que.Peek(); got != 0 || ok {
		t.Errorf("Peek() = (%v, %v), want = (%v, %v)", got, ok, 0, false)
	}
}

func TestQueueIsEmptySize(t *testing.T) {
	t.Parallel()

	que := golang.NewQueue[int]()

	que.Enqueue(1)
	que.Enqueue(2)
	que.Enqueue(3)

	if got := que.Size(); got != 3 {
		t.Errorf("Size() = %v, want = %v", got, 3)
	}

	if got := que.IsEmpty(); got {
		t.Errorf("IsEmpty() = %v, want = %v", got, false)
	}

	que.Dequeue()

	if got := que.Size(); got != 2 {
		t.Errorf("Size() = %v, want = %v", got, 2)
	}

	que.Dequeue()
	que.Dequeue()

	if got := que.Size(); got != 0 {
		t.Errorf("Size() = %v, want = %v", got, 0)
	}

	if got := que.IsEmpty(); !got {
		t.Errorf("IsEmpty() = %v, want = %v", got, true)
	}
}
