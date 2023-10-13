package counting

func findLongestSubstring(s string, k int) int {
	ans := 0
	freq := make(map[byte]int)

	for l, r := 0, 0; r < len(s); r++ {
		freq[s[r]]++

		for len(freq) > k {
			freq[s[l]]--

			if freq[s[l]] == 0 {
				delete(freq, s[l])
			}

			l++
		}

		if tmp := r - l + 1; tmp > ans {
			ans = tmp
		}
	}

	return ans
}
