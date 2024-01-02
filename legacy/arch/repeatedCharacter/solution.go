package repeatedCharacter

func repeatedCharacter(s string) byte {
	visited := make(map[byte]struct{})

	for i := range s {
		_, ok := visited[s[i]]
		if ok {
			return s[i]
		}
		visited[s[i]] = struct{}{}
	}

	// impossible due task description
	return 0
}
