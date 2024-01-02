package arrayPairSum

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)

	sum := 0
	for i, num := range nums {
		if i%2 == 0 {
			sum += num
		}
	}

	return sum
}
