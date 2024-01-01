package solutions

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	i := 0

	for j := 0; j < len(s) && i < len(g); j++ {
		if s[j] >= g[i] {
			i++
		}
	}

	return i
}
