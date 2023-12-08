package mathfunc

import "cmp"

func Max[T cmp.Ordered](nums ...T) T {
	// LeetCode use Golang < 1.21, panics below should be
	m := nums[0]
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}

func Min[T cmp.Ordered](nums ...T) T {
	// LeetCode use Golang < 1.21, panics below should be
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
