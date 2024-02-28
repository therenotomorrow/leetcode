package golang

func maxProduct(nums []int) int {
	var m1, m2 int

	for _, num := range nums {
		if num > m1 {
			m2 = m1
			m1 = num
		} else if num > m2 {
			m2 = num
		}
	}

	return (m1 - 1) * (m2 - 1)
}
