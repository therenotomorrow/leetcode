package golang

func removeInterval(intervals [][]int, toBeRemoved []int) [][]int {
	ans := make([][]int, 0)

	for _, interval := range intervals {
		if interval[0] > toBeRemoved[1] || interval[1] < toBeRemoved[0] {
			ans = append(ans, interval)

			continue
		}
		// ---|
		if interval[0] < toBeRemoved[0] {
			ans = append(ans, []int{interval[0], toBeRemoved[0]})
		}
		// |---
		if interval[1] > toBeRemoved[1] {
			ans = append(ans, []int{toBeRemoved[1], interval[1]})
		}
	}

	return ans
}
