package findMedianSortedArrays

const half = 2

// maxLeft returns the maximum element of left `nums` divided by partition `part`.
//
//	part
//
// nums: [ 1 3 8 | 9 15 ]
//
//	   ^
//	maxLeft
func maxLeft(part int, nums []int) int {
	if part == 0 {
		return -1000000
	}

	return nums[part-1]
}

// minRight returns the minimum element of right `nums` divided by partition `part`.
//
//	part
//
// nums: [ 1 3 8 | 9 15 ]
//
//	   ^
//	minRight
func minRight(part int, nums []int) int {
	if part == len(nums) {
		return 1000000
	}

	return nums[part]
}

// findMedianSortedArrays finds the median of the two sorted arraysAndStrings nums1 and nums2.
func findMedianSortedArrays(nums1, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)

	if len1 > len2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	low := 0
	high := len1

	for low <= high {
		//       low    part1   high
		// nums1: [ 1 3 8 | 9 15 ]
		//              ^   ^
		//              l1  r1
		//
		part1 := (high + low) / half

		//             part2
		// nums2: [ 7 11 | 19 21 22 25 ]
		//            ^    ^
		//            l2   r2
		//
		part2 := (len1+len2+1)/half - part1 // +1 for supporting both odds and evens lengths

		l1 := maxLeft(part1, nums1)
		l2 := maxLeft(part2, nums2)
		r1 := minRight(part1, nums1)
		r2 := minRight(part2, nums2)

		switch {
		case l1 <= r2 && l2 <= r1:
			if (len1+len2)%2 == 0 {
				return float64(Max(l1, l2)+Min(r1, r2)) / 2.0
			}

			return float64(Max(l1, l2))

		case l1 > r2:
			high = part1 - 1

		default:
			low = part1 + 1
		}
	}

	return 0 // if arraysAndStrings are not sorted (impossible situation by task description)
}

func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
