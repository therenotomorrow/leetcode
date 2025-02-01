package golang

func isArraySpecial(nums []int) bool {
	isEven := func(num int) bool {
		return num%2 == 0
	}

	for i := 1; i < len(nums); i++ {
		if isEven(nums[i-1]) == isEven(nums[i]) {
			return false
		}
	}

	return true
}
