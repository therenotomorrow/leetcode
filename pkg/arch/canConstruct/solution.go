package canConstruct

func canConstruct(ransomNote string, magazine string) bool {
	counter := make(map[rune]int)

	for _, ch := range magazine {
		counter[ch]++
	}

	for _, ch := range ransomNote {
		cnt := counter[ch]

		if cnt == 0 {
			return false
		}

		counter[ch]--
	}

	return true
}
