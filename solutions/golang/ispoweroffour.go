package golang

func isPowerOfFour(n int) bool {
	// is power of two + mod 3
	return n > 0 && n&(n-1) == 0 && n%3 == 1
}
