package golang

import "sort"

func frequencySort2(nums []int) []int {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}

	sort.SliceStable(nums, func(i, j int) bool {
		if cnt[nums[i]] == cnt[nums[j]] {
			return nums[i] > nums[j]
		}

		return cnt[nums[i]] < cnt[nums[j]]
	})

	return nums
}
