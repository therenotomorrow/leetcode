package golang

import "testing"

func TestCustomStackSmoke1(t *testing.T) {
	t.Parallel()

	obj := CustomStackConstructor(3)

	obj.Push(1)
	obj.Push(2)

	want := 2
	if got := obj.Pop(); got != want {
		t.Errorf("obj.Pop() = %v, want = %v", got, want)
	}

	obj.Push(2)
	obj.Push(3)
	obj.Push(4)

	obj.Increment(5, 100)
	obj.Increment(2, 100)

	want = 103
	if got := obj.Pop(); got != want {
		t.Errorf("obj.Pop() = %v, want = %v", got, want)
	}

	want = 202
	if got := obj.Pop(); got != want {
		t.Errorf("obj.Pop() = %v, want = %v", got, want)
	}

	want = 201
	if got := obj.Pop(); got != want {
		t.Errorf("obj.Pop() = %v, want = %v", got, want)
	}

	want = -1
	if got := obj.Pop(); got != want {
		t.Errorf("obj.Pop() = %v, want = %v", got, want)
	}
}
