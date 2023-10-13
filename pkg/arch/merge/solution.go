package merge

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := 0, 0, m

	for i < m+n && j < n {
		if nums2[j] < nums1[i] {
			for l := m + n - 1; l > i; l-- {
				nums1[l-1], nums1[l] = nums1[l], nums1[l-1]
			}
			nums1[i] = nums2[j]
			j++
			k++
		}

		i++
	}

	for j < n {
		nums1[k] = nums2[j]
		k++
		j++
	}
}
