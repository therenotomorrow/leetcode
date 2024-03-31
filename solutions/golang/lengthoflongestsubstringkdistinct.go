package golang

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	var (
		runes     = []rune(s)
		cnt       = make(map[rune]int)
		left, ans int
	)

	for right, rr := range runes {
		cnt[rr]++

		for len(cnt) > k {
			lr := runes[left]
			cnt[lr]--

			if cnt[lr] == 0 {
				delete(cnt, lr)
			}

			left++
		}

		ans = Max(ans, right-left+1)
	}

	return ans
}
