package golang

func getCommon(nums1 []int, nums2 []int) int {
	var i1, i2 int

	for i1 < len(nums1) && i2 < len(nums2) {
		switch {
		case nums1[i1] < nums2[i2]:
			i1++
		case nums1[i1] > nums2[i2]:
			i2++
		default:
			return nums1[i1]
		}
	}

	return -1
}
