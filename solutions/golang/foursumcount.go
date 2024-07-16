package golang

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	ans := 0
	cnt := make(map[int]int)

	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			cnt[num1+num2]++
		}
	}

	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			ans += cnt[-(num3 + num4)]
		}
	}

	return ans
}
