package golang

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	count := make([]PairNode, len(mat))

	for i, row := range mat {
		count[i] = PairNode{Sum(row...), i}
	}

	sort.SliceStable(count, func(i, j int) bool {
		if count[i][0] == count[j][0] {
			return i < j
		}

		return count[i][0] < count[j][0]
	})

	index := make([]int, k)
	for i, pair := range count[:k] {
		index[i] = pair[1]
	}

	return index
}
