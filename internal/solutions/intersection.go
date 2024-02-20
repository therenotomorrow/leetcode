package solutions

import "github.com/therenotomorrow/leetcode/pkg/datastruct"

func intersection(nums1 []int, nums2 []int) []int {
	set1 := datastruct.NewSet[int]()
	set2 := datastruct.NewSet[int]()

	for _, num := range nums1 {
		set1.Add(num)
	}

	for _, num := range nums2 {
		if set1.Contains(num) {
			set2.Add(num)
		}
	}

	return set2.Values()
}
