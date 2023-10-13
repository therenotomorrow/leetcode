package minSubArrayLen

func minSubArrayLen(target int, nums []int) int {
	var (
		min, sum int
		i, j     int
	)

	min = len(nums) + 1
	for ; j < len(nums); j++ {
		sum += nums[j]

		for ; sum >= target; i++ {
			sum -= nums[i]

			if j-i+1 < min {
				min = j - i + 1
			}
		}
	}

	if min > len(nums) {
		return 0
	}

	return min
}
