package golang

import "sort"

func canAttendMeetings(intervals [][]int) bool {
	const minInterval = 2

	if len(intervals) < minInterval {
		return true
	}

	sort.SliceStable(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	for i := range len(intervals) - 1 {
		if intervals[i][1] > intervals[i+1][0] {
			return false
		}
	}

	return true
}
