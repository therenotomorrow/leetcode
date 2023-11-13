package datastruct_test

import (
	"github.com/therenotomorrow/leetcode/pkg/datastruct"
	"testing"
)

func TestSetAddDelContainsInt(t *testing.T) {
	s := datastruct.NewSet[int]()

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

func TestSetPopulateInt(t *testing.T) {
	s := datastruct.NewSet[int]()

	s.Populate(1, 2, 3, 4, 5)

	for i := 1; i <= 5; i++ {
		if got := s.Contains(i); !got {
			t.Errorf("Populate() = %v, want = %v", got, i)
		}
	}
}

func TestSetAddDelContainsRune(t *testing.T) {
	s := datastruct.NewSet[rune]()

	// empty set
	if got := s.Contains('a'); got {
		t.Errorf("Contains() = %v, want = %v", got, false)
	}

	s.Add('a')
	s.Add('b')

	// regular usage `.Add()`
	if isA, isB := s.Contains('a'), s.Contains('b'); !isA || !isB {
		t.Errorf("Add() = (%v, %v), want = (%v, %v)", isA, isB, true, true)
	}

	// regular usage `.Del()`
	s.Del('a')
	if isA, isB := s.Contains('a'), s.Contains('b'); isA || !isB {
		t.Errorf("Del() = (%v, %v), want = (%v, %v)", isA, isB, false, true)
	}

	// key doesn't exist
	s.Del('c')
	if got := s.Contains('b'); !got {
		t.Errorf("Del() = %v, want = %v", got, true)
	}
}

func TestSetPopulateRune(t *testing.T) {
	s := datastruct.NewSet[rune]()

	s.Populate('a', 'b', 'c', 'd', 'e')

	for i := 0; i < 5; i++ {
		if got := s.Contains('a' + rune(i)); !got {
			t.Errorf("Populate() = %v, want = %v", got, i)
		}
	}
}
