package golang

import "math"

func closestPrimes(left int, right int) []int {
	const two = 2

	primes := make([]int, 0)
	sieveOfEratosthenes := make([]bool, right+1)

	for num := two; num <= right; num++ { // first two numbers are not primes
		sieveOfEratosthenes[num] = true
	}

	for num := two; num <= right; num++ {
		if !sieveOfEratosthenes[num] {
			continue
		}

		for multi := num * num; multi <= right; multi += num {
			sieveOfEratosthenes[multi] = false
		}
	}

	for num := left; num <= right; num++ {
		if sieveOfEratosthenes[num] {
			primes = append(primes, num)
		}
	}

	minDiff := math.MaxInt
	closest := []int{-1, -1}

	if len(primes) < two {
		return closest
	}

	for i := 1; i < len(primes); i++ {
		diff := primes[i] - primes[i-1]
		if diff < minDiff {
			minDiff = diff
			closest[0] = primes[i-1]
			closest[1] = primes[i]
		}
	}

	return closest
}
