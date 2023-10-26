package mathfunc

func Max(num int, nums ...int) int {
	// LeetCode use Golang < 1.21
	m := num
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}

func Min(num int, nums ...int) int {
	// LeetCode use Golang < 1.21
	m := num
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

func Sum(num int, nums ...int) int {
	s := num
	for _, n := range nums {
		s += n
	}
	return s
}
