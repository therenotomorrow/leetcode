package checkInclusion

func checkInclusion(s1 string, s2 string) bool {
	m1 := make(map[byte]int)
	for i := range s1 {
		m1[s1[i]]++
	}

	m2 := make(map[byte]int)
	for l, r := 0, 0; r < len(s2); r++ {
		if r >= len(s1) {
			m2[s2[l]]--
			l++
		}

		m2[s2[r]]++

		if isAnagram(m1, m2) {
			return true
		}
	}

	return false
}

func isAnagram(m1 map[byte]int, m2 map[byte]int) bool {
	for r, c := range m1 {
		if m2[r] != c {
			return false
		}
	}
	return true
}
