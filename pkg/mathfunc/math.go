package mathfunc

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
