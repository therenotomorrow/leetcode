package golang

func maxSubarrayLength(nums []int, k int) int {
	var (
		cnt       = make(map[int]int)
		left, ans int
	)

	for right, num := range nums {
		cnt[num]++

		for cnt[nums[right]] > k {
			cnt[nums[left]]--
			left++
		}

		ans = Max(ans, right-left+1)
	}

	return ans
}
