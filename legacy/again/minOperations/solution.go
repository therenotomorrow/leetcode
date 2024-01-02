package minOperations

import "sort"

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func minOperations(nums []int) int {
	n := len(nums)
	ans := n

	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}

	// sorted-unique
	idx, suNums := 0, make([]int, len(set))
	for num := range set {
		suNums[idx] = num
		idx++
	}

	sort.Ints(suNums)

	for i, j := 0, 0; i < len(suNums); i++ {
		for j < len(suNums) && suNums[j] < suNums[i]+n {
			j++
		}

		inRange := j - i
		ans = min(ans, n-inRange)
	}

	return ans
}
