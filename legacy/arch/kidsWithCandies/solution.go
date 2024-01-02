package kidsWithCandies

func max(a []int) int {
	m := a[0]
	for _, n := range a {
		if n > m {
			m = n
		}
	}

	return m
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	m := max(candies)

	for i, c := range candies {
		if c+extraCandies >= m {
			result[i] = true
		}
	}

	return result
}
