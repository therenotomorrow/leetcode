package golang

func vowelStrings(words []string, queries [][]int) []int {
	var (
		sum    = 0
		set    = NewSet[byte]()
		ans    = make([]int, len(queries))
		vowels = make([]int, len(words)+1)
	)

	set.Populate('a', 'e', 'i', 'o', 'u')

	for i, word := range words {
		if set.Contains(word[0]) && set.Contains(word[len(word)-1]) {
			sum++
		}

		vowels[i+1] = sum
	}

	for i, query := range queries {
		ans[i] = vowels[query[1]+1] - vowels[query[0]]
	}

	return ans
}
