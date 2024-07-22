package golang

import "sort"

func countPairs1(nums1 []int, nums2 []int) int64 {
	diff := make([]int, len(nums1))
	for i := range diff {
		diff[i] = nums1[i] - nums2[i]
	}

	sort.Ints(diff)

	ans := 0

	for left, right := 0, len(diff)-1; left < right; {
		if diff[left]+diff[right] > 0 {
			ans += right - left
			right--
		} else {
			left++
		}
	}

	return int64(ans)
}
