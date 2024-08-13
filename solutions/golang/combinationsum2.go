package golang

import (
	"slices"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	var (
		ans       = make([][]int, 0)
		backtrack func(curr int, rest int, part []int)
	)

	backtrack = func(curr int, rest int, part []int) {
		if rest < 0 {
			return
		}

		if rest == 0 {
			ans = append(ans, slices.Clone(part))

			return
		}

		for i := curr; i < len(candidates) && rest >= candidates[i]; i++ {
			if i > curr && candidates[i] == candidates[i-1] {
				continue
			}

			part = append(part, candidates[i])

			backtrack(i+1, rest-candidates[i], part)

			part = part[:len(part)-1]
		}
	}

	sort.Ints(candidates)

	backtrack(0, target, []int{})

	return ans
}
