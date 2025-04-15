package golang

func maxSubArrayLen(nums []int, k int) int {
	ans := 0
	prefix := 0
	cnt := make(map[int]int)

	for i, num := range nums {
		prefix += num

		if prefix == k {
			ans = i + 1
		}

		opposite := prefix - k
		if _, ok := cnt[opposite]; ok {
			ans = Max(ans, i-cnt[opposite])
		}

		if _, ok := cnt[prefix]; !ok {
			cnt[prefix] = i
		}
	}

	return ans
}
