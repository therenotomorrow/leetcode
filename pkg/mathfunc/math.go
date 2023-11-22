package mathfunc

import "math"

func Max(nums ...int) int {
	// LeetCode use Golang < 1.21
	m := math.MinInt
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}

func Min(nums ...int) int {
	// LeetCode use Golang < 1.21
	m := math.MaxInt
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
