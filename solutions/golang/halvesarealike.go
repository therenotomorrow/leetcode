package golang

func halvesAreAlike(str string) bool {
	const half = 2

	cnt := 0
	alike := map[byte]int{
		'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1,
		'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1,
	}

	for i, j := 0, len(str)/half; j < len(str); {
		cnt += alike[str[i]]
		cnt -= alike[str[j]]

		i++
		j++
	}

	return cnt == 0
}
