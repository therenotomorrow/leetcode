package golang

func reversePrefix(word string, ch byte) string {
	bytes := []byte(word)
	found := 0

	for i, b := range bytes {
		if b == ch {
			found = i

			break
		}
	}

	for i := 0; i < found; i++ {
		bytes[i], bytes[found] = bytes[found], bytes[i]
		found--
	}

	return string(bytes)
}
