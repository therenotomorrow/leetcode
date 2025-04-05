package golang

func partitionLabels(text string) []int {
	rights := make([]int, Alphabet)
	for i, r := range text {
		rights[r-'a'] = i
	}

	left := 0
	right := 0
	labels := make([]int, 0)

	for i, r := range text {
		right = Max(right, rights[r-'a'])

		if i != right {
			continue
		}

		labels = append(labels, right-left+1)
		left = right + 1
	}

	return labels
}
