package golang

func areAlmostEqual(text1 string, text2 string) bool {
	const size = 2

	var (
		cnt  int
		diff [size]int
	)

	for i := range text1 {
		if text1[i] == text2[i] {
			continue
		}

		cnt++

		switch cnt {
		case 1:
			diff[0] = i
		case size:
			diff[1] = i
		default:
			return false
		}
	}

	return text1[diff[0]] == text2[diff[1]] && text1[diff[1]] == text2[diff[0]]
}
