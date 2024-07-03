package golang

import "math"

func maxNumberOfBalloons(text string) int {
	const size = 5

	cnt := make(map[rune]int)

	for _, let := range text {
		switch let {
		case 'b', 'a', 'l', 'o', 'n':
			cnt[let]++
		}
	}

	// special case
	if len(cnt) != size {
		return 0
	}

	ans := math.MaxInt

	for let, times := range cnt {
		if let == 'l' || let == 'o' {
			times /= 2
		}

		ans = Min(ans, times)
	}

	return ans
}
