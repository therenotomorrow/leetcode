package golang

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	const limit = 100

	arr1 := make([]int, limit)
	for _, num := range nums1 {
		arr1[num] = 1
	}

	arr2 := make([]int, limit)
	for _, num := range nums2 {
		arr2[num] = 1
	}

	arr3 := make([]int, limit)
	for _, num := range nums3 {
		arr3[num] = 1
	}

	ans := make([]int, 0)

	for i := 0; i < limit; i++ {
		cnt := Sum(arr1[i], arr2[i], arr3[i])
		if cnt > 1 {
			ans = append(ans, i)
		}
	}

	return ans
}
