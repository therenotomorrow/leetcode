package golang

func maximumPopulation(logs [][]int) int {
	const (
		minYear = 1950
		maxYear = 2050
	)

	cnt := make(map[int]int)

	for year := minYear; year <= maxYear; year++ {
		for _, log := range logs {
			if log[0] <= year && year < log[1] {
				cnt[year]++
			}
		}
	}

	ans := 0
	maxPopulation := 0

	for year, population := range cnt {
		switch {
		case population > maxPopulation:
			maxPopulation = population
			ans = year
		case population == maxPopulation:
			ans = Min(year, ans)
		}
	}

	return ans
}
