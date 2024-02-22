package golang

func halvesAreAlike(s string) bool {
	cnt := 0
	alike := map[byte]int{
		'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1,
		'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1,
	}

	for i, j := 0, len(s)/2; j < len(s); {
		cnt += alike[s[i]]
		cnt -= alike[s[j]]

		i++
		j++
	}

	return cnt == 0
}
