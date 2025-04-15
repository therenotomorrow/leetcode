package golang

func getMaximumXor(nums []int, maximumBit int) []int {
	prefix := make([]int, len(nums))
	prefix[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] ^ nums[i]
	}

	ans := make([]int, len(nums))
	mask := (1 << maximumBit) - 1

	for i := range nums {
		curr := prefix[len(prefix)-i-1]
		ans[i] = curr ^ mask
	}

	return ans
}
