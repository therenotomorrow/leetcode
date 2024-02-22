package golang

import "math"

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	states := minutesToTest/minutesToDie + 1

	// if we use `math.Log2(buckets)` - we use only binary [alive or die] `states`
	pigs := math.Log2(float64(buckets)) / math.Log2(float64(states))

	return int(math.Ceil(pigs))
}
