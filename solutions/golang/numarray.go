package golang

type NumArray struct {
	prefix []int
}

func NumArrayConstructor(nums []int) NumArray {
	obj := NumArray{prefix: make([]int, 1)}

	for i, num := range nums {
		obj.prefix = append(obj.prefix, obj.prefix[i]+num)
	}

	return obj
}

func (na *NumArray) SumRange(left int, right int) int {
	return na.prefix[right+1] - na.prefix[left]
}
