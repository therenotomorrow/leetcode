package golang

import "testing"

func TestMyCircularQueueSmoke1(t *testing.T) {
	t.Parallel()

	obj := MyCircularQueueConstructor(3)

	if got := obj.EnQueue(1); !got {
		t.Errorf("obj.EnQueue() = %v, want = %v", got, true)
	}

	if got := obj.EnQueue(2); !got {
		t.Errorf("obj.EnQueue() = %v, want = %v", got, true)
	}

	if got := obj.EnQueue(3); !got {
		t.Errorf("obj.EnQueue() = %v, want = %v", got, true)
	}

	if got := obj.EnQueue(4); got {
		t.Errorf("obj.EnQueue() = %v, want = %v", got, false)
	}

	want := 3
	if got := obj.Rear(); got != want {
		t.Errorf("obj.Rear() = %v, want = %v", got, want)
	}

	if got := obj.IsFull(); !got {
		t.Errorf("obj.IsFull() = %v, want = %v", got, true)
	}

	if got := obj.DeQueue(); !got {
		t.Errorf("obj.DeQueue() = %v, want = %v", got, true)
	}

	if got := obj.EnQueue(4); !got {
		t.Errorf("obj.EnQueue() = %v, want = %v", got, true)
	}

	want = 4
	if got := obj.Rear(); got != want {
		t.Errorf("obj.Rear() = %v, want = %v", got, want)
	}
}
