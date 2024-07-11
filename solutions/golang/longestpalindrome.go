package golang

func longestPalindrome(s string) int {
	cnt := make(map[rune]int)
	for _, runa := range s {
		cnt[runa]++
	}

	length := 0
	midOdd := 0

	for _, times := range cnt {
		if times%2 != 0 {
			length-- // make every character even
			midOdd = 1
		}

		length += times
	}

	return length + midOdd
}
