package solutions

import (
	"cmp"
)

func maxLength(arr []string) int {
	var dynamic func(curr int, str string) int

	dynamic = func(currIdx int, currStr string) int {
		set := NewSet[rune]()
		for _, r := range currStr {
			set.Add(r)
		}

		maxLen := set.Size()

		if maxLen != len(currStr) {
			return 0
		}

		for i := currIdx; i < len(arr); i++ {
			maxLen = Max(maxLen, dynamic(i+1, currStr+arr[i]))
		}

		return maxLen
	}

	return dynamic(0, "")
}

type Set[T comparable] interface {
	Add(val T)
	Del(val T)
	Size() int
	Populate(vals ...T)
	Contains(val T) bool
	Values() []T
}

type set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	return &set[T]{data: make(map[T]struct{})}
}

func (s *set[T]) Add(val T) {
	s.data[val] = struct{}{}
}

func (s *set[T]) Del(val T) {
	delete(s.data, val)
}

func (s *set[T]) Size() int {
	return len(s.data)
}

func (s *set[T]) Populate(vals ...T) {
	for _, val := range vals {
		s.Add(val)
	}
}

func (s *set[T]) Contains(val T) bool {
	_, ok := s.data[val]

	return ok
}

func (s *set[T]) Values() []T {
	vals := make([]T, 0, s.Size())

	for val := range s.data {
		vals = append(vals, val)
	}

	return vals
}

func Max[T cmp.Ordered](nums ...T) T {
	m := nums[0]

	for _, n := range nums {
		if n > m {
			m = n
		}
	}

	return m
}
