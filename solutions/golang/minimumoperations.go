package golang

import "sort"

func minimumOperations(nums []int) int {
	sort.Ints(nums)

	ans := 0

	for {
		curr := 0

		for i, num := range nums {
			if num != 0 {
				curr = i

				break
			}
		}

		if curr == 0 && nums[curr] == 0 {
			break
		}

		val := nums[curr]

		for i := curr; i < len(nums); i++ {
			nums[i] -= val
		}

		ans++
	}

	return ans
}
