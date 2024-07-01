package golang

import (
	"testing"
)

func TestParkingSystemSmoke1(t *testing.T) {
	t.Parallel()

	obj := ParkingSystemConstructor(1, 1, 0)

	want := true
	if got := obj.AddCar(1); got != want {
		t.Errorf(" obj.AddCar() = %v, want = %v", got, want)
	}

	want = true
	if got := obj.AddCar(2); got != want {
		t.Errorf(" obj.AddCar() = %v, want = %v", got, want)
	}

	want = false
	if got := obj.AddCar(3); got != want {
		t.Errorf(" obj.AddCar() = %v, want = %v", got, want)
	}

	want = false
	if got := obj.AddCar(1); got != want {
		t.Errorf(" obj.AddCar() = %v, want = %v", got, want)
	}
}
