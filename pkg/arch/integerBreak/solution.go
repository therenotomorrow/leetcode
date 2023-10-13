package integerBreak

var cache = make(map[int]int)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func dynamic(n int) int {
	if val, ok := cache[n]; ok {
		return val
	}

	if n <= 3 {
		cache[n] = n
		return n
	}

	ans := n

	for i := 2; i < n; i++ {
		ans = max(ans, i*dynamic(n-i))
	}

	cache[n] = ans

	return ans
}

// fixme: solve again
func integerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}

	return dynamic(n)
}
