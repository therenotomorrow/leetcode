package minSubarray

func minSubarray(nums []int, p int) int {
	n := len(nums)
	totalSum := 0

	for _, num := range nums {
		totalSum = (totalSum + num) % p
	}

	target := totalSum % p
	if target == 0 {
		return 0
	}

	modMap := make(map[int]int)
	modMap[0] = -1
	currentSum := 0
	minLen := n

	for i := range n {
		currentSum = (currentSum + nums[i]) % p

		needed := (currentSum - target + p) % p

		_, ok := modMap[needed]
		if ok {
			minLen = min(minLen, i-modMap[needed])
		}

		modMap[currentSum] = i
	}

	if minLen == n {
		return -1
	}

	return minLen
}
