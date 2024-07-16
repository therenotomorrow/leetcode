package golang

import "testing"

func TestTwoSumSmoke1(t *testing.T) {
	t.Parallel()

	obj := TwoSumConstructor()

	obj.Add(1)
	obj.Add(3)
	obj.Add(5)

	if got := obj.Find(4); got != true {
		t.Errorf("obj.Find() = %v, want = %v", got, true)
	}

	if got := obj.Find(7); got != false {
		t.Errorf("obj.Find() = %v, want = %v", got, false)
	}
}

func TestTwoSumTest16WrongAnswer(t *testing.T) {
	t.Parallel()

	obj := TwoSumConstructor()

	obj.Add(0)
	obj.Add(0)

	if got := obj.Find(0); got != true {
		t.Errorf("obj.Find() = %v, want = %v", got, true)
	}
}

func TestTwoSumTest17WrongAnswer(t *testing.T) {
	t.Parallel()

	obj := TwoSumConstructor()

	obj.Add(0)

	if got := obj.Find(0); got != false {
		t.Errorf("obj.Find() = %v, want = %v", got, false)
	}
}
