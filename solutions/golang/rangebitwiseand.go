package golang

func rangeBitwiseAnd(left int, right int) int {
	// Brian Kernighan's algorithm: n & (n - 1) == common prefix
	for left < right {
		right &= right - 1
	}

	return right
}
