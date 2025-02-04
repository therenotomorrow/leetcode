package golang

func maxAscendingSum(nums []int) int {
	minNum := 0
	ans := 0
	curr := 0

	for _, num := range nums {
		if minNum < num {
			curr += num
			minNum = num
		} else {
			ans = Max(ans, curr)
			minNum = num
			curr = minNum
		}
	}

	return Max(ans, curr)
}
