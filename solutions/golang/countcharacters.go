package golang

func countOnce(word string, chars []int) int {
	for _, r := range word {
		if chars[r-'a'] < 1 {
			return 0
		}
		chars[r-'a']--
	}

	return len(word)
}

func countCharacters(words []string, chars string) int {
	used := make([]int, 26)
	for _, r := range chars {
		used[r-'a']++
	}

	ans := 0

	for _, word := range words {
		tmp := make([]int, 26)
		copy(tmp, used)
		ans += countOnce(word, tmp)
	}

	return ans
}
