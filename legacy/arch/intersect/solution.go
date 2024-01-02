package intersect

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums2) < len(nums1) {
		return intersect(nums2, nums1)
	}

	hSet := make(map[int]int)
	for _, num := range nums1 {
		hSet[num]++
	}

	result := make([]int, 0)
	for _, num := range nums2 {
		val, ok := hSet[num]
		if !ok {
			continue
		}

		if val > 0 {
			result = append(result, num)
			hSet[num]--
		}
	}

	return result
}
