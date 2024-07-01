package golang

import "testing"

func TestMovingAverageSmoke1(t *testing.T) {
	t.Parallel()

	obj := MovingAverageConstructor(3)

	want := 1.0
	if got := obj.Next(1); got != want {
		t.Errorf(" obj.Next() = %v, want = %v", got, want)
	}

	want = 5.5
	if got := obj.Next(10); got != want {
		t.Errorf(" obj.Next() = %v, want = %v", got, want)
	}

	obj.Next(3)

	want = 6.0
	if got := obj.Next(5); got != want {
		t.Errorf(" obj.Next() = %v, want = %v", got, want)
	}
}
