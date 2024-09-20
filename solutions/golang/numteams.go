package golang

func numTeams(rating []int) int {
	const teamSize = 3

	type compare func(a int, b int) bool

	var (
		memo    = make([][]int, len(rating)+1) // alternative of `cache.NewCache()`
		dynamic func(curr int, size int, cmp compare) int
	)

	dynamic = func(curr int, size int, cmp compare) int {
		if curr == len(rating) {
			return 0
		}

		if size == teamSize {
			return 1
		}

		if val := memo[curr][size]; val != -1 {
			return val
		}

		numOfTeams := 0

		for next := curr + 1; next < len(rating); next++ {
			if cmp(rating[next], rating[curr]) {
				numOfTeams += dynamic(next, size+1, cmp)
			}
		}

		memo[curr][size] = numOfTeams

		return numOfTeams
	}

	dynamicWithMemo := func(cmp compare) int {
		for i := range memo {
			memo[i] = make([]int, teamSize+1)

			for j := range memo[i] {
				memo[i][j] = -1
			}
		}

		numOfTeams := 0

		for curr := range rating {
			numOfTeams += dynamic(curr, 1, cmp)
		}

		return numOfTeams
	}

	lessCmp := func(a int, b int) bool { return a < b }
	moreCmp := func(a int, b int) bool { return a > b }

	return dynamicWithMemo(lessCmp) + dynamicWithMemo(moreCmp)
}
