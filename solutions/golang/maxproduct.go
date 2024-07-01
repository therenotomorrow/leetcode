package golang

func maxProduct(nums []int) int {
	var max1, max2 int

	for _, num := range nums {
		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}
	}

	return (max1 - 1) * (max2 - 1)
}
