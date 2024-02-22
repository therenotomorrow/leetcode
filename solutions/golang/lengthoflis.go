package golang

func lengthOfLIS(nums []int) int {
	cnt := make([]int, len(nums))

	for i, num := range nums {
		cnt[i]++

		for j := 0; j < i; j++ {
			if num > nums[j] {
				cnt[i] = Max(cnt[i], cnt[j]+1)
			}
		}
	}

	return Max(cnt...)
}
