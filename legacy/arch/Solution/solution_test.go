package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	var want []int

	// ---- case 1
	obj := Constructor([]int{1, 2, 3})

	want = []int{1, 2, 3}
	t.Run("", func(t *testing.T) {
		if got := obj.Shuffle(); !reflect.DeepEqual(obj.nums, want) {
			t.Errorf(" obj.Shuffle() = %v, want %v", got, want)
		}
	})

	want = []int{1, 2, 3}
	t.Run("", func(t *testing.T) {
		if got := obj.Reset(); !reflect.DeepEqual(got, want) {
			t.Errorf(" obj.Reset() = %v, want %v", got, want)
		}
	})
}
