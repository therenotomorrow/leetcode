package golang

func fib(n int) int {
	a := 0
	b := 1

	for i := 1; i <= n; i++ {
		a, b = b, a+b
	}

	return a
}
