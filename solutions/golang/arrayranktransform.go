package golang

import "sort"

func arrayRankTransform(arr []int) []int {
	index := make(map[int][]int)
	for i, num := range arr {
		index[num] = append(index[num], i)
	}

	sort.Ints(arr)

	rank := 1
	ranks := make([]int, len(arr))

	for _, num := range arr {
		if _, ok := index[num]; !ok {
			continue
		}

		for _, j := range index[num] {
			ranks[j] = rank
		}

		delete(index, num)

		rank++
	}

	return ranks
}
