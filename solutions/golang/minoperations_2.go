package golang

import "math"

func minOperations2(nums []int) int {
	const size = 3

	cnt := make(map[int]int)

	for _, num := range nums {
		cnt[num]++
	}

	ans := 0

	for _, times := range cnt {
		if times == 1 {
			return -1
		}

		ans += int(math.Ceil(float64(times) / size))
	}

	return ans
}
