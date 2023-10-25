package solutions

func lengthOfLastWord(s string) int {
	c := 0
	r := len(s) - 1

	for ; s[r] == ' '; r-- {
	}

	for ; r >= 0 && s[r] != ' '; r-- {
		c++
	}

	return c
}
