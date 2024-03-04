package golang

import "sort"

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)

	score := 0

	for left, right := 0, len(tokens)-1; left <= right; {
		if power >= tokens[left] {
			power -= tokens[left]

			score++
			left++

			continue
		}

		if left < right && score > 0 {
			power += tokens[right]

			score--
			right--

			continue
		}

		break
	}

	return score
}
