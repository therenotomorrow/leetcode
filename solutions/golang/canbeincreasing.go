package golang

func canBeIncreasing(nums []int) bool {
	skips := 0

	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			skips++

			if i > 1 && nums[i-2] >= nums[i] {
				nums[i] = nums[i-1]
			}
		}
	}

	return skips <= 1
}
