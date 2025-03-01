package golang

func applyOperations(nums []int) []int {
	for i := range len(nums) - 1 {
		if nums[i] != nums[i+1] {
			continue
		}

		nums[i] *= 2
		nums[i+1] = 0
	}

	idx := 0
	ans := make([]int, len(nums))

	for _, num := range nums {
		if num == 0 {
			continue
		}

		ans[idx] = num
		idx++
	}

	return ans
}
