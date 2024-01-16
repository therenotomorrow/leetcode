package solutions

import "testing"

func TestRandomizedSetSmoke1(t *testing.T) {
	obj := RandomizedSetConstructor()
	want := 0

	t.Run("", func(t *testing.T) {
		if got := obj.Insert(1); !got {
			t.Errorf(" obj.Insert() = %v, want = %v", got, true)
		}
	})

	t.Run("", func(t *testing.T) {
		if got := obj.Remove(2); got {
			t.Errorf(" obj.Remove() = %v, want = %v", got, false)
		}
	})

	t.Run("", func(t *testing.T) {
		if got := obj.Insert(2); !got {
			t.Errorf(" obj.Insert() = %v, want = %v", got, true)
		}
	})

	t.Run("", func(t *testing.T) {
		if got := obj.GetRandom(); got != 1 && got != 2 {
			t.Errorf(" obj.GetRandom() = %v, want = %v or %v", got, 1, 2)
		}
	})

	t.Run("", func(t *testing.T) {
		if got := obj.Remove(1); !got {
			t.Errorf(" obj.Remove() = %v, want = %v", got, true)
		}
	})

	t.Run("", func(t *testing.T) {
		if got := obj.Insert(2); got {
			t.Errorf(" obj.Insert() = %v, want = %v", got, false)
		}
	})

	want = 2

	t.Run("", func(t *testing.T) {
		if got := obj.GetRandom(); got != want {
			t.Errorf(" obj.GetRandom() = %v, want = %v", got, want)
		}
	})
}
