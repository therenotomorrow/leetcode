package golang

func largestAltitude(gain []int) int {
	var currAlt, maxAlt int

	for _, high := range gain {
		currAlt += high
		maxAlt = Max(maxAlt, currAlt)
	}

	return maxAlt
}
