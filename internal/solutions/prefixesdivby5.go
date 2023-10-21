package solutions

func prefixesDivBy5(nums []int) []bool {
	isDiv5 := make([]bool, len(nums))

	total := 0
	for i, num := range nums {
		// % 10 to fix overflow errors
		total = (2*total + num) % 10
		isDiv5[i] = total%5 == 0
	}

	return isDiv5
}
