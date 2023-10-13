package strStr

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		for j := 0; j < len(needle); j++ {
			if needle[j] != haystack[i+j] {
				break
			}

			if j == len(needle)-1 {
				return i
			}
		}
	}

	return -1
}
