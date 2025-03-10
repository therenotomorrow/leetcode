package golang

func countOfSubstrings(word string, k int) int64 { //nolint:cyclop
	const vowelsCnt = 5

	var (
		ans             = 0
		vowelsSet       = NewSet[byte]()
		vowelsMap       = make(map[byte]int)
		consonants      = 0
		consonantsSlice = make([]int, len(word))
	)

	vowelsSet.Populate('a', 'e', 'i', 'o', 'u')

	shrinkLeft := func(left int) int {
		letter := word[left]

		if vowelsSet.Contains(letter) {
			vowelsMap[letter]--
		} else {
			consonants--
		}

		if vowelsMap[letter] == 0 {
			delete(vowelsMap, letter)
		}

		return left + 1
	}
	shrinkRight := func(right int) {
		letter := word[right]

		if vowelsSet.Contains(letter) {
			vowelsMap[letter]++
		} else {
			consonants++
		}
	}

	for currConsonant, nextConsonant := len(word)-1, len(word); currConsonant >= 0; currConsonant-- {
		consonantsSlice[currConsonant] = nextConsonant

		if !vowelsSet.Contains(word[currConsonant]) {
			nextConsonant = currConsonant
		}
	}

	for left, right := 0, 0; right < len(word); right++ {
		shrinkRight(right)

		for consonants > k {
			left = shrinkLeft(left)
		}

		for left < len(word) && len(vowelsMap) == vowelsCnt && consonants == k {
			ans += consonantsSlice[right] - right
			left = shrinkLeft(left)
		}
	}

	return int64(ans)
}
