package golang

import (
	"math"
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
	const half = 2

	len1 := len(nums1)
	len2 := len(nums2)

	if len1 > len2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	for low, high := 0, len1; low <= high; {
		//       low    part1   high
		// nums1: [ 1 3 8 | 9 15 ]
		//              ^   ^
		//             ml1 mr1
		part1 := (high + low) / half
		ml1 := maxLeft(part1, nums1)
		mr1 := minRight(part1, nums1)

		//             part2
		// nums2: [ 7 11 | 19 21 22 25 ]
		//            ^    ^
		//           ml2  mr2
		part2 := (len1+len2+1)/half - part1 // +1 for both odds and evens lengths
		ml2 := maxLeft(part2, nums2)
		mr2 := minRight(part2, nums2)

		switch {
		case ml1 <= mr2 && ml2 <= mr1:
			if (len1+len2)%half == 0 {
				return float64(Max(ml1, ml2)+Min(mr1, mr2)) / half
			}

			return float64(Max(ml1, ml2))

		case ml1 > mr2:
			high = part1 - 1

		default:
			low = part1 + 1
		}
	}

	// impossible situation by task description: if arraysAndStrings are not sorted
	return 0
}
