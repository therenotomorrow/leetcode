package golang

func smallestEvenMultiple(num int) int {
	const two = 2

	if num%two == 0 {
		return num
	}

	return two * num
}
