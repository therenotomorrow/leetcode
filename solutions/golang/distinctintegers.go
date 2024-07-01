package golang

func distinctIntegers(num int) int {
	const two = 2

	if num < two {
		return 1
	}

	return num - 1
}
