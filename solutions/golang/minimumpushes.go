package golang

import "sort"

func minimumPushes(word string) int {
	const limit = 8

	cnt := make([]int, Alphabet)
	for _, runa := range word {
		cnt[runa-'a']++
	}

	sort.Stable(sort.Reverse(sort.IntSlice(cnt)))

	var ans, mul int

	for i, num := range cnt {
		mul = (i / limit) + 1
		ans += mul * num
	}

	return ans
}
