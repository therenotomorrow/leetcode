package solutions

import "sort"

func largestPerimeter(nums []int) int64 {
	sort.Ints(nums)

	sum := 0
	perimeter := -1

	for _, num := range nums {
		if num < sum {
			perimeter = num + sum
		}

		sum += num
	}

	return int64(perimeter)
}
