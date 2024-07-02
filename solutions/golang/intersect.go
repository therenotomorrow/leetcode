package golang

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums2) < len(nums1) {
		return intersect(nums2, nums1)
	}

	visited := make(map[int]int, len(nums1))
	intersect := make([]int, 0)

	for _, num := range nums1 {
		visited[num]++
	}

	for _, num := range nums2 {
		cnt, exist := visited[num]
		if !exist {
			continue
		}

		if cnt > 0 {
			intersect = append(intersect, num)
			visited[num]--
		}
	}

	return intersect
}
