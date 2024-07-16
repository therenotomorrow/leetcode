package golang

import "testing"

func TestMyHashSetSmoke1(t *testing.T) {
	t.Parallel()

	obj := MyHashSetConstructor()

	obj.Add(1)
	obj.Add(2)

	if got := obj.Contains(1); got != true {
		t.Errorf(" obj.Contains() = %v, want = %v", got, true)
	}

	if got := obj.Contains(3); got != false {
		t.Errorf(" obj.Contains() = %v, want = %v", got, false)
	}

	obj.Add(2)

	if got := obj.Contains(2); got != true {
		t.Errorf(" obj.Contains() = %v, want = %v", got, true)
	}

	obj.Remove(2)

	if got := obj.Contains(2); got != false {
		t.Errorf(" obj.Contains() = %v, want = %v", got, false)
	}
}
