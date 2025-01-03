package golang

func waysToSplitArray(nums []int) int {
	var cnt, lSum, rSum int

	total := Sum(nums...)

	for _, num := range nums[:len(nums)-1] {
		lSum += num
		rSum = total - lSum

		if lSum >= rSum {
			cnt++
		}
	}

	return cnt
}
