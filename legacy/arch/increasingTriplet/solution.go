package increasingTriplet

import "math"

func increasingTriplet(nums []int) bool {
	a := math.MaxInt
	b := math.MaxInt

	for _, num := range nums {
		if num <= a {
			a = num
		} else if num <= b {
			b = num
		} else {
			return true
		}
	}

	return false
}
