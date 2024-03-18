package golang

func equalSubstring(s string, t string, maxCost int) int {
	var cost, maxLen, left int

	for right := 0; right < len(s); right++ {
		cost += Abs(int(s[right]) - int(t[right]))

		for cost > maxCost {
			cost -= Abs(int(s[left]) - int(t[left]))
			left++
		}

		maxLen = Max(maxLen, right-left+1)
	}

	return maxLen
}
