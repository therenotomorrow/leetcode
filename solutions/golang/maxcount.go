package golang

func maxCount(banned []int, n int, maxSum int) int {
	cnt := 0
	ban := NewSet[int]()

	ban.Populate(banned...)

	for i := 1; i <= n; i++ {
		if ban.Contains(i) {
			continue
		}

		if maxSum-i < 0 {
			break
		}

		cnt++
		maxSum -= i
	}

	return cnt
}
