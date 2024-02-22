package golang

import "sort"

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)

	dynamic := make([][]int, len(nums))

	for i, num := range nums {
		maxSet := make([]int, 0)

		for j := 0; j < i; j++ {
			if num%nums[j] == 0 {
				if len(dynamic[j]) > len(maxSet) {
					maxSet = make([]int, len(dynamic[j]))

					copy(maxSet, dynamic[j])
				}
			}
		}

		maxSet = append(maxSet, num) //nolint:makezero

		dynamic[i] = maxSet
	}

	sort.SliceStable(dynamic, func(i, j int) bool {
		return len(dynamic[i]) > len(dynamic[j])
	})

	return dynamic[0]
}
