package golang

import "sort"

func minMeetingRooms(intervals [][]int) int {
	start := make([]int, len(intervals))
	end := make([]int, len(intervals))

	for i, interval := range intervals {
		start[i] = interval[0]
		end[i] = interval[1]
	}

	sort.Ints(start)
	sort.Ints(end)

	var startPnt, endPnt, rooms int

	for startPnt < len(intervals) {
		// we could start the meeting
		if start[startPnt] >= end[endPnt] {
			endPnt++
			rooms--
		}

		startPnt++
		rooms++
	}

	return rooms
}
