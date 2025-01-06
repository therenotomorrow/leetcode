package golang

func minOperations5(boxes string) []int {
	moves := make([]int, len(boxes))

	for i, box := range boxes {
		if box != '1' {
			continue
		}

		for j := range boxes {
			moves[j] += Abs(i - j)
		}
	}

	return moves
}
