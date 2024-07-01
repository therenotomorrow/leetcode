package golang

func sumOfDigits(nums []int) int {
	minNum := Min(nums...)

	sum := 0

	for minNum > 0 {
		sum += minNum % Digits
		minNum /= Digits
	}

	return sum%2 ^ 1
}
