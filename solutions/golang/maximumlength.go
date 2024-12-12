package golang

func maximumLength(text string) int {
	const thrice = 3

	ans := -1

	for length := 1; length <= len(text); length++ {
		count := make(map[string]int)

		for i := range len(text) - length + 1 {
			substr := text[i : i+length]
			found := true

			for j := range substr {
				if substr[j] != substr[0] {
					found = false

					break
				}
			}

			if found {
				count[substr]++
			}
		}

		for _, freq := range count {
			if freq >= thrice {
				ans = length
			}
		}
	}

	return ans
}
