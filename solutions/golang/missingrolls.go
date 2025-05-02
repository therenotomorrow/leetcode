package golang

func missingRolls(rolls []int, mean int, n int) []int {
	missingSum := mean*(len(rolls)+n) - Sum(rolls...)

	if n > missingSum || missingSum > 6*n {
		return []int{}
	}

	mean = missingSum / n
	remainder := missingSum % n
	missing := make([]int, n)

	for i := range n {
		missing[i] = mean
	}

	for i := range remainder {
		missing[i]++
	}

	return missing
}
