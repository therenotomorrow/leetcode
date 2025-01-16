package golang

func xorAllNums(nums1 []int, nums2 []int) int {
	xor1 := 0
	xor2 := 0

	if len(nums2)%2 != 0 {
		for _, num := range nums1 {
			xor1 ^= num
		}
	}

	if len(nums1)%2 != 0 {
		for _, num := range nums2 {
			xor2 ^= num
		}
	}

	return xor1 ^ xor2
}
