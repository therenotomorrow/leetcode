package maximumBeauty

import (
	"sort"
)

func maximumBeauty(items [][]int, queries []int) []int {
	ans := make([]int, len(queries))

	sort.SliceStable(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	queriesWithIndices := make([][2]int, len(queries))
	for i := range queries {
		queriesWithIndices[i][0] = queries[i]
		queriesWithIndices[i][1] = i
	}

	sort.Slice(queriesWithIndices, func(i, j int) bool {
		return queriesWithIndices[i][0] < queriesWithIndices[j][0]
	})

	itemIndex := 0
	maxBeauty := 0

	for i := range queries {
		query := queriesWithIndices[i][0]
		originalIndex := queriesWithIndices[i][1]

		for itemIndex < len(items) && items[itemIndex][0] <= query {
			maxBeauty = max(maxBeauty, items[itemIndex][1])
			itemIndex++
		}

		ans[originalIndex] = maxBeauty
	}

	return ans
}
