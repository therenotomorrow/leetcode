package solutions

import "github.com/therenotomorrow/leetcode/pkg/mathfunc"

func containsNearbyDuplicate(nums []int, k int) bool {
	cnt := make(map[int]int)

	for i, num := range nums {
		j, ok := cnt[num]

		if ok && mathfunc.Abs(i-j) <= k {
			return true
		}

		cnt[num] = i
	}

	return false
}
