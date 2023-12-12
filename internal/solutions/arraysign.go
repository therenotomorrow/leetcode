package solutions

func signFunc(x int) int {
	switch {
	case x > 0:
		return 1
	case x < 0:
		return -1
	default:
		return 0
	}
}

func arraySign(nums []int) int {
	prod := 1

	for _, num := range nums {
		prod *= signFunc(num)
	}

	return prod
}
