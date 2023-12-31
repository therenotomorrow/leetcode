package solutions

import "github.com/therenotomorrow/leetcode/internal/structs"

func countNicePairs(nums []int) int {
	diff := make([]int, len(nums))

	for i, num := range nums {
		// took from reverse.go: https://leetcode.com/problems/reverse-integer/
		diff[i] = num - reverse(num)
	}

	cnt := 0
	seen := make(map[int]int)

	for _, num := range diff {
		cnt = (cnt + seen[num]) % structs.MOD
		seen[num]++
	}

	return cnt
}
