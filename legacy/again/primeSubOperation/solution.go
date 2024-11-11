package primeSubOperation

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeSubOperation(nums []int) bool {
	// Find the maximum element in nums
	maxElement := nums[0]
	for _, num := range nums {
		if num > maxElement {
			maxElement = num
		}
	}

	// Generate the largest prime less than or equal to each number up to maxElement
	previousPrime := make([]int, maxElement+1)
	for i := 2; i <= maxElement; i++ {
		if isPrime(i) {
			previousPrime[i] = i
		} else {
			previousPrime[i] = previousPrime[i-1]
		}
	}

	// Perform the subtraction operation on nums
	for i := 0; i < len(nums); i++ {
		bound := nums[i]
		if i > 0 {
			bound -= nums[i-1]
		}

		if bound <= 0 {
			return false
		}

		// Subtract the largest prime less than the current bound
		nums[i] -= previousPrime[bound-1]
	}
	return true
}
