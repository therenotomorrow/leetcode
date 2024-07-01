package golang

import "sort"

func findContentChildren(gArr []int, sArr []int) int {
	sort.Ints(gArr)
	sort.Ints(sArr)

	i := 0

	for j := 0; j < len(sArr) && i < len(gArr); j++ {
		if sArr[j] >= gArr[i] {
			i++
		}
	}

	return i
}
