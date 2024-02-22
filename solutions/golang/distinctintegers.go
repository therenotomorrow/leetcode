package golang

func distinctIntegers(n int) int {
	if n < 2 {
		return 1
	}

	return n - 1
}
