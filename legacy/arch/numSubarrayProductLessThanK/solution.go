package numSubarrayProductLessThanK

func numSubarrayProductLessThanK(nums []int, k int) int {
	var (
		curr = 1
		res  = 0
	)

	for left, right := 0, 0; right < len(nums); right++ {
		curr *= nums[right]

		for ; curr >= k && left < len(nums); left++ {
			curr /= nums[left]
		}

		res += right - left + 1
	}

	if res < 0 {
		res = 0
	}

	return res
}
