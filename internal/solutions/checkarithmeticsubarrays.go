package solutions

import (
	"github.com/therenotomorrow/leetcode/pkg/datastruct"
	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
)

func isArithmetic(arr []int) bool {
	minNum := mathfunc.Min(arr...)
	maxNum := mathfunc.Max(arr...)

	if (maxNum-minNum)%(len(arr)-1) != 0 {
		// because the diff is not the int value
		return false
	}

	diff := (maxNum - minNum) / (len(arr) - 1)

	set := datastruct.NewSet[int]()
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

	for i := 0; i < size; i++ {
		bools[i] = isArithmetic(nums[l[i] : r[i]+1])
	}

	return bools
}
