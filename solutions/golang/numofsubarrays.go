package golang

func numOfSubarrays(arr []int) int {
	var cnt, prefix, odds, evens int

	evens = 1 // because 0 is the even number

	for _, num := range arr {
		prefix += num

		if prefix%2 == 0 {
			cnt += odds
			evens++
		} else {
			cnt += evens
			odds++
		}
	}

	return cnt % MOD
}
