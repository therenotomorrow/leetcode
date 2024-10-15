package golang

import "testing"

func TestFirstUniqueSmoke1(t *testing.T) {
	t.Parallel()

	obj := FirstUniqueConstructor([]int{2, 3, 5})

	want := 2
	if got := obj.ShowFirstUnique(); got != want {
		t.Errorf("obj.ShowFirstUnique() = %v, want = %v", got, want)
	}

	obj.Add(5)

	want = 2
	if got := obj.ShowFirstUnique(); got != want {
		t.Errorf("obj.ShowFirstUnique() = %v, want = %v", got, want)
	}

	obj.Add(2)

	want = 3
	if got := obj.ShowFirstUnique(); got != want {
		t.Errorf("obj.ShowFirstUnique() = %v, want = %v", got, want)
	}

	obj.Add(3)

	want = -1
	if got := obj.ShowFirstUnique(); got != want {
		t.Errorf("obj.ShowFirstUnique() = %v, want = %v", got, want)
	}
}

func TestFirstUniqueSmoke2(t *testing.T) {
	t.Parallel()

	obj := FirstUniqueConstructor([]int{7, 7, 7, 7, 7, 7})

	want := -1
	if got := obj.ShowFirstUnique(); got != want {
		t.Errorf("obj.ShowFirstUnique() = %v, want = %v", got, want)
	}

	obj.Add(7)
	obj.Add(3)
	obj.Add(3)
	obj.Add(7)
	obj.Add(17)

	want = 17
	if got := obj.ShowFirstUnique(); got != want {
		t.Errorf("obj.ShowFirstUnique() = %v, want = %v", got, want)
	}
}

func TestFirstUniqueSmoke3(t *testing.T) {
	t.Parallel()

	obj := FirstUniqueConstructor([]int{809})

	want := 809
	if got := obj.ShowFirstUnique(); got != want {
		t.Errorf("obj.ShowFirstUnique() = %v, want = %v", got, want)
	}

	obj.Add(809)

	want = -1
	if got := obj.ShowFirstUnique(); got != want {
		t.Errorf("obj.ShowFirstUnique() = %v, want = %v", got, want)
	}
}
