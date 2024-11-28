package golang

import (
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var (
		cnt  = make(map[int]int)
		ans  = make([]int, 0, len(arr1))
		rest = make([]int, 0)
	)

	for _, num := range arr1 {
		cnt[num]++
	}

	for _, num := range arr2 {
		for cnt[num] > 0 {
			ans = append(ans, num)
			cnt[num]--
		}

		cnt[num] = -1
	}

	for _, num := range arr1 {
		if cnt[num] == -1 {
			continue
		}

		rest = append(rest, num)
	}

	sort.Ints(rest)

	return append(ans, rest...)
}
