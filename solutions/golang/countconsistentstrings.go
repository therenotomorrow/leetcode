package golang

func countConsistentStrings(allowed string, words []string) int {
	cnt := 0
	mask := make([]int, Alphabet)

	for _, runa := range allowed {
		mask[int(runa-'a')]++
	}

	for _, word := range words {
		found := true

		for _, runa := range word {
			if mask[int(runa-'a')] == 0 {
				found = false

				break
			}
		}

		if !found {
			continue
		}

		cnt++
	}

	return cnt
}
