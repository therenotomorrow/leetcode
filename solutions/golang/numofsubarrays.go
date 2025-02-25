package golang

func numOfSubarrays(arr []int) int {
	var cnt, prefixSum, odds, evens int

	evens = 1 // because 0 is the even number

	for _, num := range arr {
		prefixSum += num

		if prefixSum%2 == 0 {
			cnt += odds
			evens++
		} else {
			cnt += evens
			odds++
		}
	}

	return cnt % MOD
}
