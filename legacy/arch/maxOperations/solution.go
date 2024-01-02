package maxOperations

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)

	var (
		res = 0
		i   = 0
		j   = len(nums) - 1
	)

	for i < j {
		switch val := nums[i] + nums[j]; {
		case val < k:
			i++
		case val > k:
			j--
		default:
			i++
			j--
			res++
		}
	}

	return res
}
