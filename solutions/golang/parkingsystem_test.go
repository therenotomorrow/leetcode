package golang

import (
	"testing"
)

func TestParkingSystemSmoke1(t *testing.T) {
	t.Parallel()

	obj := ParkingSystemConstructor(1, 1, 0)

	if got := obj.AddCar(1); got != true {
		t.Errorf("obj.AddCar() = %v, want = %v", got, true)
	}

	if got := obj.AddCar(2); got != true {
		t.Errorf("obj.AddCar() = %v, want = %v", got, true)
	}

	if got := obj.AddCar(3); got != false {
		t.Errorf("obj.AddCar() = %v, want = %v", got, false)
	}

	if got := obj.AddCar(1); got != false {
		t.Errorf("obj.AddCar() = %v, want = %v", got, false)
	}
}
