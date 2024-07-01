package golang

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Sort(sort.Reverse(sort.IntSlice(happiness)))

	ans := 0

	for cycle := range k {
		happy := happiness[cycle] - cycle
		if happy < 0 {
			happy = 0
		}

		ans += happy
	}

	return int64(ans)
}
