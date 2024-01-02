package datastruct_test

import (
	"testing"

	"github.com/therenotomorrow/leetcode/pkg/datastruct"
)

func TestQueueEnqueueDequeue(t *testing.T) {
	queue := datastruct.NewQueue[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	if got, ok := queue.Dequeue(); got != 1 || !ok {
		t.Errorf("Dequeue() = (%v, %v), want = (%v, %v)", got, ok, 1, true)
	}

	if got, ok := queue.Dequeue(); got != 2 || !ok {
		t.Errorf("Dequeue() = (%v, %v), want = (%v, %v)", got, ok, 2, true)
	}

	if got, ok := queue.Dequeue(); got != 3 || !ok {
		t.Errorf("Dequeue() = (%v, %v), want = (%v, %v)", got, ok, 3, true)
	}

	if got, ok := queue.Dequeue(); got != 0 || ok {
		t.Errorf("Dequeue() = (%v, %v), want = (%v, %v)", got, ok, 0, false)
	}
}

func TestQueuePeek(t *testing.T) {
	queue := datastruct.NewQueue[int]()

	queue.Enqueue(42)

	if got, ok := queue.Peek(); got != 42 || !ok {
		t.Errorf("Peek() = (%v, %v), want = (%v, %v)", got, ok, 42, true)
	}

	queue.Dequeue()

	if got, ok := queue.Peek(); got != 0 || ok {
		t.Errorf("Peek() = (%v, %v), want = (%v, %v)", got, ok, 0, false)
	}
}

func TestQueueIsEmptySize(t *testing.T) {
	queue := datastruct.NewQueue[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	if got := queue.Size(); got != 3 {
		t.Errorf("Size() = %v, want = %v", got, 3)
	}

	if got := queue.IsEmpty(); got {
		t.Errorf("IsEmpty() = %v, want = %v", got, false)
	}

	queue.Dequeue()

	if got := queue.Size(); got != 2 {
		t.Errorf("Size() = %v, want = %v", got, 2)
	}

	queue.Dequeue()
	queue.Dequeue()

	if got := queue.Size(); got != 0 {
		t.Errorf("Size() = %v, want = %v", got, 0)
	}

	if got := queue.IsEmpty(); !got {
		t.Errorf("IsEmpty() = %v, want = %v", got, true)
	}
}
