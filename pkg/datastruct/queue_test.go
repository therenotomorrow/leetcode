package datastruct_test

import (
	"github.com/therenotomorrow/leetcode/pkg/datastruct"
	"testing"
)

func TestQueueEnqueueDequeue(t *testing.T) {
	q := datastruct.NewQueue[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if got, ok := q.Dequeue(); got != 1 || !ok {
		t.Errorf("Dequeue() = (%v, %v), want = (%v, %v)", got, ok, 1, true)
	}

	if got, ok := q.Dequeue(); got != 2 || !ok {
		t.Errorf("Dequeue() = (%v, %v), want = (%v, %v)", got, ok, 2, true)
	}

	if got, ok := q.Dequeue(); got != 3 || !ok {
		t.Errorf("Dequeue() = (%v, %v), want = (%v, %v)", got, ok, 3, true)
	}

	if got, ok := q.Dequeue(); got != 0 || ok {
		t.Errorf("Dequeue() = (%v, %v), want = (%v, %v)", got, ok, 0, false)
	}
}

func TestQueuePeek(t *testing.T) {
	q := datastruct.NewQueue[int]()

	q.Enqueue(42)

	if got, ok := q.Peek(); got != 42 || !ok {
		t.Errorf("Peek() = (%v, %v), want = (%v, %v)", got, ok, 42, true)
	}

	q.Dequeue()

	if got, ok := q.Peek(); got != 0 || ok {
		t.Errorf("Peek() = (%v, %v), want = (%v, %v)", got, ok, 0, false)
	}
}

func TestQueueIsEmptySize(t *testing.T) {
	s := datastruct.NewQueue[int]()

	s.Enqueue(1)
	s.Enqueue(2)
	s.Enqueue(3)

	if got := s.Size(); got != 3 {
		t.Errorf("Size() = %v, want = %v", got, 3)
	}
	if got := s.IsEmpty(); got {
		t.Errorf("IsEmpty() = %v, want = %v", got, false)
	}

	s.Dequeue()

	if got := s.Size(); got != 2 {
		t.Errorf("Size() = %v, want = %v", got, 2)
	}

	s.Dequeue()
	s.Dequeue()

	if got := s.Size(); got != 0 {
		t.Errorf("Size() = %v, want = %v", got, 0)
	}
	if got := s.IsEmpty(); !got {
		t.Errorf("IsEmpty() = %v, want = %v", got, true)
	}
}
