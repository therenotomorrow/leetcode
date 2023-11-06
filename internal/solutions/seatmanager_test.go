package solutions

import (
	"reflect"
	"testing"
)

func TestSeatManagerSmoke1(t *testing.T) {
	obj := Constructor(5)
	want := 0

	want = 1
	t.Run("", func(t *testing.T) {
		if got := obj.Reserve(); got != want {
			t.Errorf(" obj.Reserve() = %v, want = %v", got, want)
		}
	})

	want = 2
	t.Run("", func(t *testing.T) {
		if got := obj.Reserve(); got != want {
			t.Errorf(" obj.Reserve() = %v, want = %v", got, want)
		}
	})

	obj.Unreserve(2)

	want = 2
	t.Run("", func(t *testing.T) {
		if got := obj.Reserve(); got != want {
			t.Errorf(" obj.Reserve() = %v, want = %v", got, want)
		}
	})

	want = 3
	t.Run("", func(t *testing.T) {
		if got := obj.Reserve(); got != want {
			t.Errorf(" obj.Reserve() = %v, want = %v", got, want)
		}
	})

	want = 4
	t.Run("", func(t *testing.T) {
		if got := obj.Reserve(); got != want {
			t.Errorf(" obj.Reserve() = %v, want = %v", got, want)
		}
	})

	want = 5
	t.Run("", func(t *testing.T) {
		if got := obj.Reserve(); got != want {
			t.Errorf(" obj.Reserve() = %v, want = %v", got, want)
		}
	})

	obj.Unreserve(5)
}

func TestSeatManagerWrongAnswer13(t *testing.T) {
	obj := Constructor(798)
	got := make([]int, 0)

	got = append(got, obj.Reserve())
	obj.Unreserve(1)
	got = append(got, obj.Reserve())
	obj.Unreserve(1)
	got = append(got, obj.Reserve())
	obj.Unreserve(1)
	got = append(got, obj.Reserve())
	got = append(got, obj.Reserve())
	got = append(got, obj.Reserve())
	got = append(got, obj.Reserve())
	got = append(got, obj.Reserve())
	obj.Unreserve(5)
	obj.Unreserve(3)
	obj.Unreserve(2)
	got = append(got, obj.Reserve())
	obj.Unreserve(4)
	got = append(got, obj.Reserve())
	got = append(got, obj.Reserve())
	obj.Unreserve(1)
	obj.Unreserve(3)
	got = append(got, obj.Reserve())
	obj.Unreserve(2)
	obj.Unreserve(4)
	obj.Unreserve(1)
	got = append(got, obj.Reserve())
	obj.Unreserve(1)
	got = append(got, obj.Reserve())
	got = append(got, obj.Reserve())
	got = append(got, obj.Reserve())
	got = append(got, obj.Reserve())
	obj.Unreserve(1)
	got = append(got, obj.Reserve())
	got = append(got, obj.Reserve())
	got = append(got, obj.Reserve())
	obj.Unreserve(1)
	obj.Unreserve(3)
	obj.Unreserve(2)
	got = append(got, obj.Reserve())
	obj.Unreserve(5)
	got = append(got, obj.Reserve())
	got = append(got, obj.Reserve())
	got = append(got, obj.Reserve())
	obj.Unreserve(5)
	got = append(got, obj.Reserve())
	obj.Unreserve(6)
	obj.Unreserve(4)
	obj.Unreserve(5)
	obj.Unreserve(1)
	got = append(got, obj.Reserve())
	obj.Unreserve(2)
	obj.Unreserve(3)
	got = append(got, obj.Reserve())
	obj.Unreserve(1)
	obj.Unreserve(2)
	got = append(got, obj.Reserve())
	got = append(got, obj.Reserve())
	got = append(got, obj.Reserve())
	got = append(got, obj.Reserve())
	obj.Unreserve(2)
	got = append(got, obj.Reserve())
	//[1 1 1 1 2 3 4 5 2 3 4 1 1 1 2 3 4 1 2 3 1 2 3 4 5 1 2 1 2 3 4 2]
	//[1 1 1 1 2 3 4 5 2 3 4 1 1 1 2 3 4 1 5 6 1 2 3 5 5 1 2 1 2 3 4 2]
	want := []int{1, 1, 1, 1, 2, 3, 4, 5, 2, 3, 4, 1, 1, 1, 2, 3, 4, 1, 5, 6, 1, 2, 3, 5, 5, 1, 2, 1, 2, 3, 4, 2}
	t.Run("", func(t *testing.T) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf(" obj.Reserve() = %v, want = %v", got, want)
		}
	})
}
