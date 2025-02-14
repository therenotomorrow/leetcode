package golang

import "testing"

func TestProductOfNumbersSmoke1(t *testing.T) {
	t.Parallel()

	obj := ProductOfNumbersConstructor()

	obj.Add(3)
	obj.Add(0)
	obj.Add(2)
	obj.Add(5)
	obj.Add(4)

	want := 20
	if got := obj.GetProduct(2); got != want {
		t.Errorf("obj.GetProduct() = %v, want = %v", got, want)
	}

	want = 40
	if got := obj.GetProduct(3); got != want {
		t.Errorf("obj.GetProduct() = %v, want = %v", got, want)
	}

	want = 0
	if got := obj.GetProduct(4); got != want {
		t.Errorf("obj.GetProduct() = %v, want = %v", got, want)
	}

	obj.Add(8)

	want = 32
	if got := obj.GetProduct(2); got != want {
		t.Errorf("obj.GetProduct() = %v, want = %v", got, want)
	}
}
