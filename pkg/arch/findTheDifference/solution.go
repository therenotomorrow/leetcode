package findTheDifference

func findTheDifference(s string, t string) byte {
	alphabet := [26]int{}

	for i := 0; i < len(s); i++ {
		alphabet[s[i]-'a']++
		alphabet[t[i]-'a']--
	}

	alphabet[t[len(s)]-'a']--

	for i, cnt := range alphabet {
		if cnt == -1 {
			return byte(i) + 'a'
		}
	}

	// impossible by text description
	return 0
}
