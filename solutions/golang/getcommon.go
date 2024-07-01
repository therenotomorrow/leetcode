package golang

func getCommon(nums1 []int, nums2 []int) int {
	var idx1, idx2 int

	for idx1 < len(nums1) && idx2 < len(nums2) {
		switch {
		case nums1[idx1] < nums2[idx2]:
			idx1++
		case nums1[idx1] > nums2[idx2]:
			idx2++
		default:
			return nums1[idx1]
		}
	}

	return -1
}
