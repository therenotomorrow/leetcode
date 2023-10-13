package lengthOfLongestSubstring

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLongestSubstring(s string) int {
	used := make(map[byte]int)

	var last, ans int
	for i := range s {
		last = max(last, used[s[i]])
		ans = max(ans, i-last+1)
		used[s[i]] = i + 1
	}

	return ans
}
