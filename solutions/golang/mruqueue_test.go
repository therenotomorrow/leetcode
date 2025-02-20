package golang

import "testing"

func TestMRUQueueSmoke1(t *testing.T) {
	t.Parallel()

	obj := MRUQueueConstructor(8)

	want := 3
	if got := obj.Fetch(3); got != want {
		t.Errorf("obj.Fetch() = %v, want = %v", got, want)
	}

	want = 6
	if got := obj.Fetch(5); got != want {
		t.Errorf("obj.Fetch() = %v, want = %v", got, want)
	}

	want = 2
	if got := obj.Fetch(2); got != want {
		t.Errorf("obj.Fetch() = %v, want = %v", got, want)
	}

	want = 2
	if got := obj.Fetch(8); got != want {
		t.Errorf("obj.Fetch() = %v, want = %v", got, want)
	}
}
