package datastruct_test

import (
	"github.com/therenotomorrow/leetcode/pkg/datastruct"
	"testing"
)

func TestSetAddDelContains(t *testing.T) {
	s := datastruct.NewSet()

	// empty set
	if got := s.Contains(42); got {
		t.Errorf("Contains() = %v, want = %v", got, false)
	}

	s.Add(42)
	s.Add(36)

	// regular usage `.Add()`
	if is42, is36 := s.Contains(42), s.Contains(36); !is42 || !is36 {
		t.Errorf("Add() = (%v, %v), want = (%v, %v)", is42, is36, true, true)
	}

	// regular usage `.Del()`
	s.Del(42)
	if is42, is36 := s.Contains(42), s.Contains(36); is42 || !is36 {
		t.Errorf("Del() = (%v, %v), want = (%v, %v)", is42, is36, false, true)
	}

	// key doesn't exist
	s.Del(99)
	if got := s.Contains(36); !got {
		t.Errorf("Del() = %v, want = %v", got, true)
	}
}

func TestSetPopulate(t *testing.T) {
	s := datastruct.NewSet()

	s.Populate(1, 2, 3, 4, 5)

	for i := 1; i <= 5; i++ {
		if got := s.Contains(i); !got {
			t.Errorf("Populate() = %v, want = %v", got, i)
		}
	}
}
