package golang

import "testing"

func TestMyHashMapSmoke1(t *testing.T) {
	t.Parallel()

	obj := MyHashMapConstructor()

	obj.Put(1, 1)
	obj.Put(2, 2)

	want := 1
	if got := obj.Get(1); got != want {
		t.Errorf("obj.Get() = %v, want = %v", got, want)
	}

	want = -1
	if got := obj.Get(3); got != want {
		t.Errorf("obj.Get() = %v, want = %v", got, want)
	}

	obj.Put(2, 1)

	want = 1
	if got := obj.Get(2); got != want {
		t.Errorf("obj.Get() = %v, want = %v", got, want)
	}

	obj.Remove(2)

	want = -1
	if got := obj.Get(2); got != want {
		t.Errorf("obj.Get() = %v, want = %v", got, want)
	}
}
