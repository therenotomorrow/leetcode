package golang

func longestNiceSubarray(nums []int) int {
	var left, track, ans int

	for right, num := range nums {
		for track&num != 0 {
			track ^= nums[left]
			left++
		}

		track |= num

		ans = Max(ans, right-left+1)
	}

	return ans
}
