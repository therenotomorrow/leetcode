package golang

func maximumSum(nums []int) int {
	cnt := make(map[int]int)
	ans := -1
	sum := func(num int) int {
		sum := 0

		for num > 0 {
			sum += num % Digits
			num /= Digits
		}

		return sum
	}

	for _, num := range nums {
		key := sum(num)
		val := cnt[key]

		if val != 0 {
			ans = Max(ans, val+num)
		}

		cnt[key] = Max(val, num)
	}

	return ans
}
