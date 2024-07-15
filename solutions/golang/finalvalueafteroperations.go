package golang

func finalValueAfterOperations(operations []string) int {
	num := 0

	for _, operation := range operations {
		switch operation[1] {
		case '+':
			num++
		case '-':
			num--
		}
	}

	return num
}
