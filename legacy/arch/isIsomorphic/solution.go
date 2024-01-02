package isIsomorphic

func isIsomorphic(s string, t string) bool {
	ms := make(map[byte]byte)
	mt := make(map[byte]byte)

	for i := range s {
		_, ok1 := ms[s[i]]
		_, ok2 := mt[t[i]]

		if !ok1 {
			ms[s[i]] = t[i]
		}
		if !ok2 {
			mt[t[i]] = s[i]
		}

		if ms[s[i]] != t[i] || mt[t[i]] != s[i] {
			return false
		}
	}

	return true
}
