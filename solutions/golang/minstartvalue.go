package golang

func minStartValue(nums []int) int {
	var minVal, curSum int

	for _, num := range nums {
		curSum += num
		minVal = Min(minVal, curSum)
	}

	return Abs(minVal) + 1
}
