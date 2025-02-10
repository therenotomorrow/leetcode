package golang

import "testing"

func TestNumberContainersSmoke1(t *testing.T) {
	t.Parallel()

	obj := NumberContainersConstructor()

	want := -1
	if got := obj.Find(10); got != want {
		t.Errorf("obj.Find() = %v, want = %v", got, want)
	}

	obj.Change(2, 10)
	obj.Change(1, 10)
	obj.Change(3, 10)
	obj.Change(5, 10)

	want = 1
	if got := obj.Find(10); got != want {
		t.Errorf("obj.Find() = %v, want = %v", got, want)
	}

	obj.Change(1, 20)

	want = 2
	if got := obj.Find(10); got != want {
		t.Errorf("obj.Find() = %v, want = %v", got, want)
	}
}
