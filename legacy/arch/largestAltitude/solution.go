package largestAltitude

func largestAltitude(gain []int) int {
	max := 0
	start := 0

	for _, high := range gain {
		start += high
		if start > max {
			max = start
		}
	}

	return max
}
