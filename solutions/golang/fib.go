package golang

func fib(n int) int {
	prev := 0
	curr := 1

	for i := 1; i <= n; i++ {
		prev, curr = curr, prev+curr
	}

	return prev
}
