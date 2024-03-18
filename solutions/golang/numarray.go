package golang

type NumArray struct {
	prefixes []int
}

func NumArrayConstructor(nums []int) NumArray {
	obj := NumArray{prefixes: make([]int, 1)}

	for i, num := range nums {
		obj.prefixes = append(obj.prefixes, obj.prefixes[i]+num)
	}

	return obj
}

func (na *NumArray) SumRange(left int, right int) int {
	return na.prefixes[right+1] - na.prefixes[left]
}
