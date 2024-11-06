package golang

func canSortArray(nums []int) bool {
	for i := range len(nums) - 1 {
		for j := range len(nums) - i - 1 {
			if nums[j] <= nums[j+1] {
				continue
			}

			// took from sortbybits.go: https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/description/
			if countBits(nums[j]) == countBits(nums[j+1]) {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			} else {
				return false
			}
		}
	}

	return true
}
