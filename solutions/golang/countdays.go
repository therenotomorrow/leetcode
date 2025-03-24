package golang

import "sort"

func countDays(days int, meetings [][]int) int {
	var ans, latest int

	sort.SliceStable(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	for _, meeting := range meetings {
		start := meeting[0]
		end := meeting[1]

		if start > latest+1 {
			ans += start - latest - 1
		}

		latest = Max(latest, end)
	}

	return ans + days - latest
}
