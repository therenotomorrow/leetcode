package equalSubstring

func abs(a byte, b byte) int {
	res := int(a) - int(b)

	if res < 0 {
		res = -res
	}

	return res
}

func equalSubstring(s string, t string, maxCost int) int {
	var (
		cost, res   int
		left, right int
	)

	for ; right < len(s); right++ {
		cost += abs(s[right], t[right])

		for cost > maxCost {
			cost -= abs(s[left], t[left])
			left++
		}

		if right-left+1 > res {
			res = right - left + 1
		}
	}

	return res
}
