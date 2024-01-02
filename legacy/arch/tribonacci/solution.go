package tribonacci

func tribonacci(n int) int {
	fib := [3]int{0, 1, 1}

	for ; n > 2; n-- {
		fib[2], fib[1], fib[0] = fib[2]+fib[1]+fib[0], fib[2], fib[1]
	}

	return fib[n]
}
