package golang

import (
	"testing"
)

func TestParkingSystemSmoke1(t *testing.T) {
	obj := ParkingSystemConstructor(1, 1, 0)
	want := false

	want = true

	t.Run("", func(t *testing.T) {
		if got := obj.AddCar(1); got != want {
			t.Errorf(" obj.AddCar() = %v, want = %v", got, want)
		}
	})

	want = true

	t.Run("", func(t *testing.T) {
		if got := obj.AddCar(2); got != want {
			t.Errorf(" obj.AddCar() = %v, want = %v", got, want)
		}
	})

	want = false

	t.Run("", func(t *testing.T) {
		if got := obj.AddCar(3); got != want {
			t.Errorf(" obj.AddCar() = %v, want = %v", got, want)
		}
	})

	want = false

	t.Run("", func(t *testing.T) {
		if got := obj.AddCar(1); got != want {
			t.Errorf(" obj.AddCar() = %v, want = %v", got, want)
		}
	})
}
