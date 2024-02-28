package golang

func findMaxLength(nums []int) int {
	var (
		ans, sum int
		count    = make(map[int]int)
	)

	for i, num := range nums {
		if num == 1 {
			sum++
		} else {
			sum--
		}

		if sum == 0 {
			ans = Max(ans, i+1)
		}

		if val, ok := count[sum]; ok {
			ans = Max(ans, i-val)
		} else {
			count[sum] = i
		}
	}

	return ans
}
