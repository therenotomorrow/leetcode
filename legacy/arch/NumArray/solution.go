package NumArray

type NumArray struct {
	prefixes []int
}

func Constructor(nums []int) NumArray {
	na := NumArray{prefixes: make([]int, len(nums))}

	na.prefixes[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		na.prefixes[i] = na.prefixes[i-1] + nums[i]
	}

	return na
}

func (na *NumArray) SumRange(left int, right int) int {
	if left < 1 {
		return na.prefixes[right]
	}

	return na.prefixes[right] - na.prefixes[left-1]
}
