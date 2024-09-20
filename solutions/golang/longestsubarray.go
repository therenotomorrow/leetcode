package golang

func longestSubarray(nums []int) int {
	ans := 0
	cnt := 0
	maxVal := Max(nums...)

	for _, num := range nums {
		if num == maxVal {
			cnt++
		} else {
			cnt = 0
		}

		ans = Max(ans, cnt)
	}

	return ans
}
