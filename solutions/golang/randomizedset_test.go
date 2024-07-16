package golang

import "testing"

func TestRandomizedSetSmoke1(t *testing.T) {
	t.Parallel()

	obj := RandomizedSetConstructor()

	if got := obj.Insert(1); !got {
		t.Errorf("obj.Insert() = %v, want = %v", got, true)
	}

	if got := obj.Remove(2); got {
		t.Errorf("obj.Remove() = %v, want = %v", got, false)
	}

	if got := obj.Insert(2); !got {
		t.Errorf("obj.Insert() = %v, want = %v", got, true)
	}

	if got := obj.GetRandom(); got != 1 && got != 2 {
		t.Errorf("obj.GetRandom() = %v, want = %v or %v", got, 1, 2)
	}

	if got := obj.Remove(1); !got {
		t.Errorf("obj.Remove() = %v, want = %v", got, true)
	}

	if got := obj.Insert(2); got {
		t.Errorf("obj.Insert() = %v, want = %v", got, false)
	}

	want := 2
	if got := obj.GetRandom(); got != want {
		t.Errorf("obj.GetRandom() = %v, want = %v", got, want)
	}
}
