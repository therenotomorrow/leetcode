package golang

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	indexed := make([][2]int, len(score))

	for i, s := range score {
		indexed[i] = [2]int{s, i}
	}

	sort.SliceStable(indexed, func(i, j int) bool {
		return indexed[i][0] > indexed[j][0]
	})

	ans := make([]string, len(score))
	top := 0

	for _, pair := range indexed {
		pos := pair[1]

		switch top {
		case 0:
			ans[pos] = "Gold Medal"
		case 1:
			ans[pos] = "Silver Medal"
		case 2:
			ans[pos] = "Bronze Medal"
		default:
			ans[pos] = strconv.Itoa(top + 1)
		}

		top++
	}

	return ans
}
