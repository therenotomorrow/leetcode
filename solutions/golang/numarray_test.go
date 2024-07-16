package golang

import "testing"

func TestNumArraySmoke1(t *testing.T) {
	t.Parallel()

	obj := NumArrayConstructor([]int{-2, 0, 3, -5, 2, -1})

	want := 1
	if got := obj.SumRange(0, 2); got != want {
		t.Errorf("obj.SumRange() = %v, want = %v", got, want)
	}

	want = -1
	if got := obj.SumRange(2, 5); got != want {
		t.Errorf("obj.SumRange() = %v, want = %v", got, want)
	}

	want = -3
	if got := obj.SumRange(0, 5); got != want {
		t.Errorf("obj.SumRange() = %v, want = %v", got, want)
	}
}
