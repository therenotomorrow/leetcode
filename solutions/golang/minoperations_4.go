package golang

func minOperations4(nums []int) int {
	cnt := 0
	curr := nums[0]

	for _, num := range nums[1:] {
		diff := 0

		if num <= curr {
			diff = curr - num + 1
			cnt += diff
		}

		curr = num + diff
	}

	return cnt
}
