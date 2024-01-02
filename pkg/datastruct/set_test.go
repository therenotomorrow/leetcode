package datastruct_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/therenotomorrow/leetcode/pkg/datastruct"
)

func TestSetAddDelLenContains(t *testing.T) {
	set := datastruct.NewSet[int]()

	// empty set
	if got := set.Contains(42); got {
		t.Errorf("Contains() = %v, want = %v", got, false)
	}

	set.Add(42)
	set.Add(36)

	// regular usage `.Add()`
	if is42, is36 := set.Contains(42), set.Contains(36); !is42 || !is36 {
		t.Errorf("Add() = (%v, %v), want = (%v, %v)", is42, is36, true, true)
	}

	if got, want := set.Len(), 2; got != want {
		t.Errorf("Len() = %v, want = %v", got, want)
	}

	// regular usage `.Del()`
	set.Del(42)

	if is42, is36 := set.Contains(42), set.Contains(36); is42 || !is36 {
		t.Errorf("Del() = (%v, %v), want = (%v, %v)", is42, is36, false, true)
	}

	if got, want := set.Len(), 1; got != want {
		t.Errorf("Len() = %v, want = %v", got, want)
	}

	// key doesn't exist
	set.Del(99)

	if got := set.Contains(36); !got {
		t.Errorf("Del() = %v, want = %v", got, true)
	}

	if got, want := set.Len(), 1; got != want {
		t.Errorf("Len() = %v, want = %v", got, want)
	}
}

func TestSetLenPopulateValues(t *testing.T) {
	set := datastruct.NewSet[int]()

	set.Populate(1, 2, 3, 4, 5)

	for i := 1; i <= 5; i++ {
		if got := set.Contains(i); !got {
			t.Errorf("Populate() = %v, want = %v", got, i)
		}
	}

	got := set.Values()
	sort.Ints(got)

	want := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values() = %v, want = %v", got, want)
	}

	if set.Len() != 5 {
		t.Errorf("Len() = %v, want = %v", set.Len(), 5)
	}
}
