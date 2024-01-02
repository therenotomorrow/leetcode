package moveZeroes

func moveZeroes(nums []int) {
	var i, j int

	for ; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}

	for i = j; i < len(nums); i++ {
		nums[i] = 0
	}
}
