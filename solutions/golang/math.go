package golang

import "cmp"

// Max returns maximum of `nums` of type `T` and have more friendly interface then build-in `max()`.
// Panics if `nums` have zero elements.
func Max[T cmp.Ordered](nums ...T) T {
	m := nums[0]

	for _, n := range nums {
		if n > m {
			m = n
		}
	}

	return m
}

// Min returns minimum of `nums` of type `T` and have more friendly interface then build-in `min()`.
// Panics if `nums` have zero elements.
func Min[T cmp.Ordered](nums ...T) T {
	m := nums[0]

	for _, n := range nums {
		if n < m {
			m = n
		}
	}

	return m
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func Sum(nums ...int) int {
	s := 0

	for _, n := range nums {
		s += n
	}

	return s
}

func Manhattan(x1 int, y1 int, x2 int, y2 int) int {
	return Abs(x1-x2) + Abs(y1-y2)
}
