package golang

import "testing"

func TestValidWordAbbrSmoke1(t *testing.T) {
	t.Parallel()

	obj := ValidWordAbbrConstructor([]string{"deer", "door", "cake", "card"})

	if got := obj.IsUnique("dear"); got != false {
		t.Errorf("obj.IsUnique() = %v, want = %v", got, false)
	}

	if got := obj.IsUnique("cart"); got != true {
		t.Errorf("obj.IsUnique() = %v, want = %v", got, true)
	}

	if got := obj.IsUnique("cane"); got != false {
		t.Errorf("obj.IsUnique() = %v, want = %v", got, false)
	}

	if got := obj.IsUnique("make"); got != true {
		t.Errorf("obj.IsUnique() = %v, want = %v", got, true)
	}

	if got := obj.IsUnique("cake"); got != true {
		t.Errorf("obj.IsUnique() = %v, want = %v", got, true)
	}
}

func TestValidWordAbbrTest46WrongAnswer(t *testing.T) {
	t.Parallel()

	obj := ValidWordAbbrConstructor([]string{"deer", "door", "cake", "card"})

	if got := obj.IsUnique("dear"); got != false {
		t.Errorf("obj.IsUnique() = %v, want = %v", got, false)
	}

	if got := obj.IsUnique("door"); got != false {
		t.Errorf("obj.IsUnique() = %v, want = %v", got, false)
	}

	if got := obj.IsUnique("cart"); got != true {
		t.Errorf("obj.IsUnique() = %v, want = %v", got, true)
	}

	if got := obj.IsUnique("cake"); got != true {
		t.Errorf("obj.IsUnique() = %v, want = %v", got, true)
	}
}
