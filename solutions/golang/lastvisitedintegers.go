package golang

func lastVisitedIntegers(nums []int) []int {
	ans := make([]int, 0)
	seen := make([]int, 0)
	cnt := 0

	for _, num := range nums {
		if num != -1 {
			cnt = 0

			seen = append([]int{num}, seen...)

			continue
		}

		if cnt >= len(seen) {
			ans = append(ans, num)
		} else {
			ans = append(ans, seen[cnt])
		}

		cnt++
	}

	return ans
}
