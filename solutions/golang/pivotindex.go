package golang

func pivotIndex(nums []int) int {
	lSum := 0
	rSum := Sum(nums...)

	for i, num := range nums {
		rSum -= num

		if lSum == rSum {
			return i
		}

		lSum += num
	}

	return -1
}
