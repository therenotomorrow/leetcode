package golang

func getMaximumXor(nums []int, maximumBit int) []int {
	prefixes := make([]int, len(nums))
	prefixes[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		prefixes[i] = prefixes[i-1] ^ nums[i]
	}

	ans := make([]int, len(nums))
	mask := (1 << maximumBit) - 1

	for i := range nums {
		curr := prefixes[len(prefixes)-i-1]
		ans[i] = curr ^ mask
	}

	return ans
}
