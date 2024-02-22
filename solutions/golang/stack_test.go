package golang_test

import (
	"testing"

	"github.com/therenotomorrow/leetcode/solutions/golang"
)

func TestStackPushPop(t *testing.T) {
	stack := golang.NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if got, ok := stack.Pop(); got != 3 || !ok {
		t.Errorf("Pop() = (%v, %v), want = (%v, %v)", got, ok, 3, true)
	}

	if got, ok := stack.Pop(); got != 2 || !ok {
		t.Errorf("Pop() = (%v, %v), want = (%v, %v)", got, ok, 2, true)
	}

	if got, ok := stack.Pop(); got != 1 || !ok {
		t.Errorf("Pop() = (%v, %v), want = (%v, %v)", got, ok, 1, true)
	}

	if got, ok := stack.Pop(); got != 0 || ok {
		t.Errorf("Pop() = (%v, %v), want = (%v, %v)", got, ok, 0, false)
	}
}

func TestStackPeek(t *testing.T) {
	stack := golang.NewStack[int]()

	stack.Push(42)

	if got, ok := stack.Peek(); got != 42 || !ok {
		t.Errorf("Peek() = (%v, %v), want = (%v, %v)", got, ok, 42, true)
	}

	stack.Pop()

	if got, ok := stack.Peek(); got != 0 || ok {
		t.Errorf("Peek() = (%v, %v), want = (%v, %v)", got, ok, 0, false)
	}
}

func TestStackIsEmptySize(t *testing.T) {
	stack := golang.NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if got := stack.Size(); got != 3 {
		t.Errorf("Size() = %v, want = %v", got, 3)
	}

	if got := stack.IsEmpty(); got {
		t.Errorf("IsEmpty() = %v, want = %v", got, false)
	}

	stack.Pop()

	if got := stack.Size(); got != 2 {
		t.Errorf("Size() = %v, want = %v", got, 2)
	}

	stack.Pop()
	stack.Pop()

	if got := stack.Size(); got != 0 {
		t.Errorf("Size() = %v, want = %v", got, 0)
	}

	if got := stack.IsEmpty(); !got {
		t.Errorf("IsEmpty() = %v, want = %v", got, true)
	}
}
