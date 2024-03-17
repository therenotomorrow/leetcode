package golang

func insert(intervals [][]int, newInterval []int) [][]int {
	i := 0
	ans := make([][]int, 0)

	// pass less intervals
	for ; i < len(intervals) && intervals[i][1] < newInterval[0]; i++ {
		ans = append(ans, intervals[i])
	}

	// try to find new start and end for newInterval
	for ; i < len(intervals) && newInterval[1] >= intervals[i][0]; i++ {
		newInterval[0] = Min(newInterval[0], intervals[i][0])
		newInterval[1] = Max(newInterval[1], intervals[i][1])
	}

	ans = append(ans, newInterval)

	for ; i < len(intervals); i++ {
		ans = append(ans, intervals[i])
	}

	return ans
}
