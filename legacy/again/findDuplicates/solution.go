package findDuplicates

import "sort"

func findDuplicates(nums []int) []int {
	ans := make([]int, 0)

	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			ans = append(ans, nums[i])
			i++
		}
	}

	return ans
}
