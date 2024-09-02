package golang

func chalkReplacer(chalk []int, k int) int {
	i := 0

	k %= Sum(chalk...)

	for ; chalk[i] <= k; i++ {
		k -= chalk[i]
	}

	return i
}
