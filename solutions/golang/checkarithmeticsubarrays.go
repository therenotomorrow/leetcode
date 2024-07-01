package golang

func isArithmetic(arr []int) bool {
	minNum := Min(arr...)
	maxNum := Max(arr...)

	if (maxNum-minNum)%(len(arr)-1) != 0 {
		// because the diff is not the int value
		return false
	}

	diff := (maxNum - minNum) / (len(arr) - 1)

	set := NewSet[int]()
	set.Populate(arr...)

	for i := range arr {
		if !set.Contains(minNum + diff*i) {
			return false
		}
	}

	return true
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	size := len(l)
	bools := make([]bool, size)

	for i := range size {
		bools[i] = isArithmetic(nums[l[i] : r[i]+1])
	}

	return bools
}
