package golang

func truncateSentence(text string, k int) string {
	cnt := 0

	for i, runa := range text {
		if runa == ' ' {
			cnt++
		}

		if cnt == k {
			return text[:i]
		}
	}

	return text
}
