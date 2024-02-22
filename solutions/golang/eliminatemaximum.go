package golang

import "sort"

func eliminateMaximum(dist []int, speed []int) int {
	times := make([]int, len(dist))

	for i := range times {
		// ceil time because of discrete of unit
		times[i] = (dist[i] + speed[i] - 1) / speed[i]
	}

	sort.Ints(times)

	cnt := 0

	for i, time := range times {
		if time <= i {
			break
		}

		cnt++
	}

	return cnt
}
