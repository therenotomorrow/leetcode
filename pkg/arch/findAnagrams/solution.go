package findAnagrams

func findAnagrams(s string, p string) []int {
	ans := make([]int, 0)

	mp := make(map[byte]int)
	for i := range p {
		mp[p[i]]++
	}

	ms := make(map[byte]int)
	for l, r := 0, 0; r < len(s); r++ {
		if r >= len(p) {
			ms[s[l]]--
			l++
		}

		ms[s[r]]++

		if isAnagram(mp, ms) {
			ans = append(ans, l)
		}
	}

	return ans
}

func isAnagram(m1 map[byte]int, m2 map[byte]int) bool {
	for r, c := range m1 {
		if m2[r] != c {
			return false
		}
	}
	return true
}
