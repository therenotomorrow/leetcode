package solutions

import (
	"math"

	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
)

func maxLeft(part int, nums []int) int {
	if part == 0 {
		return math.MinInt
	}
	// nums: [ 1 3 8 | 9 15 ], maxLeft: 8
	return nums[part-1]
}

func minRight(part int, nums []int) int {
	if part == len(nums) {
		return math.MaxInt
	}

	// nums: [ 1 3 8 | 9 15 ], minRight: 9
	return nums[part]
}

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)

	if len1 > len2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	for low, high := 0, len1; low <= high; {
		//       low    part1   high
		// nums1: [ 1 3 8 | 9 15 ]
		//              ^   ^
		//              l1  r1
		part1 := (high + low) / 2
		l1 := maxLeft(part1, nums1)
		r1 := minRight(part1, nums1)

		//             part2
		// nums2: [ 7 11 | 19 21 22 25 ]
		//            ^    ^
		//            l2   r2
		part2 := (len1+len2+1)/2 - part1 // +1 for both odds and evens lengths
		l2 := maxLeft(part2, nums2)
		r2 := minRight(part2, nums2)

		switch {
		case l1 <= r2 && l2 <= r1:
			if (len1+len2)%2 == 0 {
				return float64(mathfunc.Max(l1, l2)+mathfunc.Min(r1, r2)) / 2
			}

			return float64(mathfunc.Max(l1, l2))

		case l1 > r2:
			high = part1 - 1

		default:
			low = part1 + 1
		}
	}

	// impossible situation by task description: if arraysAndStrings are not sorted
	return 0
}
