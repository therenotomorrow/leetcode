package golang

func prefixesDivBy5(nums []int) []bool {
	const div = 5

	isDiv5 := make([]bool, len(nums))

	total := 0
	for i, num := range nums {
		// % `Digits` to fix overflow errors
		total = (2*total + num) % Digits
		isDiv5[i] = total%div == 0
	}

	return isDiv5
}
