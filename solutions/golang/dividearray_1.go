package golang

import "sort"

func divideArray1(nums []int, k int) [][]int {
	const size = 3

	ans := make([][]int, len(nums)/size)

	sort.Ints(nums)

	for i, j := 0, 0; i < len(nums); i += size {
		if nums[i+2]-nums[i] > k {
			return [][]int{}
		}

		ans[j] = nums[i : i+size]
		j++
	}

	return ans
}
