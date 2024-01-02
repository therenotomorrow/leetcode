package countPrimes

import "math"

func countPrimes(n int) int {
	// Sieve of Eratosthenes
	nums := make([]bool, n)
	maxP := int(math.Sqrt(float64(n)))

	for p := 2; p <= maxP; p++ {
		for i := p * p; i < n; i += p {
			nums[i] = true
		}
	}

	cnt := 0
	for i := 2; i < n; i++ {
		if !nums[i] {
			cnt++
		}
	}

	return cnt
}
