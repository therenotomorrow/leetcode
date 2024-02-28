package golang

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))

	i := 0
	j := len(nums) - 1

	for k := len(nums) - 1; i <= j; k-- {
		iPow := nums[i] * nums[i]
		jPow := nums[j] * nums[j]

		if iPow > jPow {
			res[k] = iPow
			i++
		} else {
			res[k] = jPow
			j--
		}
	}

	return res
}
