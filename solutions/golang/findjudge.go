package golang

func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return n // only one who we could trust
	}

	judges := make(map[int]int)

	for _, persons := range trust {
		judges[persons[0]]--
		judges[persons[1]]++
	}

	for judge, trusted := range judges {
		if trusted == n-1 {
			return judge
		}
	}

	return -1
}
