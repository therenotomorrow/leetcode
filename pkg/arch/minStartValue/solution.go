package minStartValue

func minStartValue(nums []int) int {
	min := 0
	sum := 0
	for _, num := range nums {
		sum += num
		if sum < min {
			min = sum
		}
	}

	if min < 0 {
		min = -min
	}

	return min + 1
}
