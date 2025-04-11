package golang

func minOperations8(nums []int, k int) int {
	swap := make(map[int]bool)

	for _, num := range nums {
		if num < k {
			return -1
		}

		if num > k {
			swap[num] = true
		}
	}

	return len(swap)
}
