package golang

func pivotArray(nums []int, pivot int) []int {
	ans := make([]int, len(nums))

	var less, equals, greater int

	for _, num := range nums {
		switch {
		case num < pivot:
			less++
		case num == pivot:
			equals++
		}
	}

	greater = less + equals
	equals = less
	less = 0

	for _, num := range nums {
		switch {
		case num < pivot:
			ans[less] = num
			less++
		case num > pivot:
			ans[greater] = num
			greater++
		case num == pivot:
			ans[equals] = num
			equals++
		}
	}

	return ans
}
