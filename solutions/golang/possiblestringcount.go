package golang

func possibleStringCount(word string) int {
	ans := 1 // initial string counts also

	for i := 1; i < len(word); i++ {
		if word[i-1] == word[i] {
			ans++
		}
	}

	return ans
}
