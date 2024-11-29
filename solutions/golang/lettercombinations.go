package golang

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var (
		letters = map[byte]string{
			'2': "abc", '3': "def",
			'4': "ghi", '5': "jkl",
			'6': "mno", '7': "pqrs",
			'8': "tuv", '9': "wxyz",
		}
		combinations = make([]string, 0)
		backtrack    func(curr int, combination []rune)
	)

	backtrack = func(curr int, combination []rune) {
		if len(combination) == len(digits) {
			combinations = append(combinations, string(combination))

			return
		}

		for _, letter := range letters[digits[curr]] {
			combination = append(combination, letter)
			backtrack(curr+1, combination)
			combination = combination[:len(combination)-1]
		}
	}

	backtrack(0, []rune{})

	return combinations
}
