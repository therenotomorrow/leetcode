package golang

import "sort"

func divideArray(nums []int, k int) [][]int {
	ans := make([][]int, len(nums)/3)

	sort.Ints(nums)

	for i, j := 0, 0; i < len(nums); i += 3 {
		if nums[i+2]-nums[i] > k {
			return [][]int{}
		}

		ans[j] = nums[i : i+3]
		j++
	}

	return ans
}
