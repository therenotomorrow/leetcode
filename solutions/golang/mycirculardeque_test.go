package golang

import "testing"

func TestMyCircularDequeSmoke1(t *testing.T) {
	t.Parallel()

	obj := MyCircularDequeConstructor(3)

	if got := obj.InsertLast(1); !got {
		t.Errorf("obj.InsertLast() = %v, want = %v", got, true)
	}

	if got := obj.InsertLast(2); !got {
		t.Errorf("obj.InsertLast() = %v, want = %v", got, true)
	}

	if got := obj.InsertFront(3); !got {
		t.Errorf("obj.InsertFront() = %v, want = %v", got, true)
	}

	if got := obj.InsertFront(4); got {
		t.Errorf("obj.InsertFront() = %v, want = %v", got, false)
	}

	want := 2
	if got := obj.GetRear(); got != want {
		t.Errorf("obj.GetRear() = %v, want = %v", got, want)
	}

	if got := obj.IsFull(); !got {
		t.Errorf("obj.IsFull() = %v, want = %v", got, true)
	}

	if got := obj.DeleteLast(); !got {
		t.Errorf("obj.DeleteLast() = %v, want = %v", got, true)
	}

	if got := obj.InsertFront(4); !got {
		t.Errorf("obj.InsertFront() = %v, want = %v", got, true)
	}

	want = 4
	if got := obj.GetFront(); got != want {
		t.Errorf("obj.GetFront() = %v, want = %v", got, want)
	}
}
