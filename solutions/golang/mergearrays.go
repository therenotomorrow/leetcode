package golang

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	i := 0
	j := 0
	ans := make([][]int, 0)

	for i < len(nums1) && j < len(nums2) {
		var val []int

		switch id1, id2 := nums1[i][0], nums2[j][0]; {
		case id1 == id2:
			val = []int{id1, nums1[i][1] + nums2[j][1]}
			i++
			j++
		case id1 < id2:
			val = nums1[i]
			i++
		case id2 < id1:
			val = nums2[j]
			j++
		}

		ans = append(ans, val)
	}

	for ; i < len(nums1); i++ {
		ans = append(ans, nums1[i])
	}

	for ; j < len(nums2); j++ {
		ans = append(ans, nums2[j])
	}

	return ans
}
