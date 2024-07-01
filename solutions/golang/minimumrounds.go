package golang

import "math"

func minimumRounds(tasks []int) int {
	const size = 3

	cnt := make(map[int]int)

	for _, task := range tasks {
		cnt[task]++
	}

	ans := 0

	for _, times := range cnt {
		if times == 1 {
			return -1
		}

		ans += int(math.Ceil(float64(times) / size))
	}

	return ans
}
