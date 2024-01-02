package sortedSquares

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))

	i := 0
	j := len(nums) - 1

	for k := len(nums) - 1; i <= j; k-- {
		if abs(nums[i]) > abs(nums[j]) {
			res[k] = nums[i] * nums[i]
			i++
		} else {
			res[k] = nums[j] * nums[j]
			j--
		}
	}

	return res
}
