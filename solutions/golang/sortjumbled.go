package golang

import "sort"

func sortJumbled(mapping []int, nums []int) []int {
	transform := func(n int) int {
		m := 0
		where := 1

		if n == 0 {
			return mapping[n]
		}

		for n > 0 {
			m += where * mapping[n%Digits]

			where *= Digits
			n /= Digits
		}

		return m
	}

	sort.SliceStable(nums, func(i, j int) bool {
		return transform(nums[i]) < transform(nums[j])
	})

	return nums
}
