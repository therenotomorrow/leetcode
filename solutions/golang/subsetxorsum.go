package golang

func subsetXORSum(nums []int) int {
	var backtrack func(curr int, sum int) int

	backtrack = func(curr int, sum int) int {
		if curr == len(nums) {
			return sum
		}

		return backtrack(curr+1, sum^nums[curr]) + backtrack(curr+1, sum)
	}

	return backtrack(0, 0)
}
