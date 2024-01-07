package numberOfArithmeticSlices

import "math"

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	ans := 0
	cnt := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		cnt[i] = make(map[int]int)

		for j := 0; j < i; j++ {
			delta := nums[i] - nums[j]

			if delta < math.MinInt || delta > math.MaxInt {
				continue
			}

			diff := delta
			sum := cnt[j][diff]
			origin := cnt[i][diff]

			cnt[i][diff] = origin + sum + 1

			ans += sum
		}
	}

	return ans
}
