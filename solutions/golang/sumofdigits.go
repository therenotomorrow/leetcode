package golang

func sumOfDigits(nums []int) int {
	minNum := Min(nums...)

	sum := 0

	for minNum > 0 {
		sum += minNum % 10
		minNum /= 10
	}

	return sum%2 ^ 1
}
