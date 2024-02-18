package mostBooked

import (
	"math"
	"sort"
)

func mostBooked(n int, meetings [][]int) int {
	roomAvailabilityTime := make([]int, n)
	meetingCount := make([]int, n)

	sort.SliceStable(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	for _, meeting := range meetings {
		start := meeting[0]
		end := meeting[1]
		minRoomAvailabilityTime := math.MaxInt
		minAvailableTimeRoom := 0
		foundUnusedRoom := false

		for i := 0; i < n; i++ {
			if roomAvailabilityTime[i] <= start {
				foundUnusedRoom = true
				meetingCount[i]++
				roomAvailabilityTime[i] = end
				break
			}

			if minRoomAvailabilityTime > roomAvailabilityTime[i] {
				minRoomAvailabilityTime = roomAvailabilityTime[i]
				minAvailableTimeRoom = i
			}
		}

		if !foundUnusedRoom {
			roomAvailabilityTime[minAvailableTimeRoom] += end - start
			meetingCount[minAvailableTimeRoom]++
		}
	}

	maxMeetingCount := 0
	maxMeetingCountRoom := 0

	for i := 0; i < n; i++ {
		if meetingCount[i] > maxMeetingCount {
			maxMeetingCount = meetingCount[i]
			maxMeetingCountRoom = i
		}
	}

	return maxMeetingCountRoom
}
