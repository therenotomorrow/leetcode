package solutions

func climbStairs(n int) int {
	// will start Fibonacci sequence from 1, 1, 2, ...
	onFirst, onSecond := 1, 1

	for i := 2; i <= n; i++ {
		onFirst, onSecond = onSecond, onFirst+onSecond
	}

	return onSecond
}
