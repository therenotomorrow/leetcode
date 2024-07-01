package golang

func productExceptSelf(nums []int) []int {
	const magicTwo = 2

	lProd := make([]int, len(nums))

	lProd[0] = 1
	for i := 1; i < len(nums); i++ {
		lProd[i] = lProd[i-1] * nums[i-1]
	}

	rProd := make([]int, len(nums))

	rProd[len(nums)-1] = 1
	for i := len(nums) - magicTwo; i >= 0; i-- {
		rProd[i] = rProd[i+1] * nums[i+1]
	}

	for i := range nums {
		nums[i] = lProd[i] * rProd[i]
	}

	return nums
}
