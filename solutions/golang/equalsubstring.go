package golang

func equalSubstring(str string, tpl string, maxCost int) int {
	var cost, maxLen, left int

	for right := range len(str) {
		cost += Abs(int(str[right]) - int(tpl[right]))

		for cost > maxCost {
			cost -= Abs(int(str[left]) - int(tpl[left]))
			left++
		}

		maxLen = Max(maxLen, right-left+1)
	}

	return maxLen
}
