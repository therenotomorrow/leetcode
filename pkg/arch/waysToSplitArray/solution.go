package waysToSplitArray

func waysToSplitArray(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	var lSum, rSum, cnt int
	for _, num := range nums[:len(nums)-1] {
		lSum += num
		rSum = total - lSum
		if lSum >= rSum {
			cnt++
		}
	}

	return cnt
}
