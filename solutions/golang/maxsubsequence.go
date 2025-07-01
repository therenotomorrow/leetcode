package golang

import "sort"

func maxSubsequence(nums []int, k int) []int {
	indexes := make([]PairNode, len(nums))

	for i, num := range nums {
		indexes[i] = PairNode{i, num}
	}

	sort.SliceStable(indexes, func(i, j int) bool {
		return indexes[i][1] > indexes[j][1]
	})

	indexes = indexes[:k]

	sort.SliceStable(indexes, func(i, j int) bool {
		return indexes[i][0] < indexes[j][0]
	})

	ans := make([]int, k)

	for idx, num := range indexes {
		ans[idx] = num[1]
	}

	return ans
}
