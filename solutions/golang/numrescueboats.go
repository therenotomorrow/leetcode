package golang

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	ans := 0

	for i, j := 0, len(people)-1; i <= j; j-- {
		ans++

		if people[i]+people[j] <= limit {
			i++
		}
	}

	return ans
}
