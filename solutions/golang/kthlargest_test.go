package golang

import "testing"

func TestKthLargestSmoke1(t *testing.T) {
	t.Parallel()

	obj := KthLargestConstructor(3, []int{4, 5, 8, 2})

	want := 4
	if got := obj.Add(3); got != want {
		t.Errorf("obj.Add() = %v, want = %v", got, want)
	}

	want = 5
	if got := obj.Add(5); got != want {
		t.Errorf("obj.Add() = %v, want = %v", got, want)
	}

	want = 5
	if got := obj.Add(10); got != want {
		t.Errorf("obj.Add() = %v, want = %v", got, want)
	}

	want = 8
	if got := obj.Add(9); got != want {
		t.Errorf("obj.Add() = %v, want = %v", got, want)
	}

	want = 8
	if got := obj.Add(4); got != want {
		t.Errorf("obj.Add() = %v, want = %v", got, want)
	}
}
