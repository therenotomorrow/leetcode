package numberOfWays

const MOD int = 1e9 + 7

func numberOfWays(corridor string) int {
	indexes := make([]int, 0)
	for i, item := range corridor {
		if item == 'S' {
			indexes = append(indexes, i)
		}
	}

	if len(indexes) == 0 || len(indexes)%2 == 1 {
		return 0
	}

	cnt := 1
	prev := 1
	curr := 2

	for curr < len(indexes) {
		cnt *= indexes[curr] - indexes[prev]
		cnt %= MOD
		prev += 2
		curr += 2
	}

	return cnt
}
