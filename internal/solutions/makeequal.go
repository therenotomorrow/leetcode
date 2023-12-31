package solutions

func makeEqual(words []string) bool {
	cnt := [26]int{}

	for _, word := range words {
		for _, r := range word {
			cnt[r-'a']++
		}
	}

	for _, times := range cnt {
		if times%len(words) != 0 {
			return false
		}
	}

	return true
}
