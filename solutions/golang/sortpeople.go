package golang

import "sort"

func sortPeople(names []string, heights []int) []string {
	heightMap := make(map[int]string, len(heights))

	for i, height := range heights {
		heightMap[height] = names[i]
	}

	sort.Sort(sort.Reverse(sort.IntSlice(heights)))

	for i, height := range heights {
		names[i] = heightMap[height]
	}

	return names
}
