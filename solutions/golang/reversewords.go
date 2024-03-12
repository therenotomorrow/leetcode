package golang

func reverseWords(s string) string {
	space := -1
	bytes := []byte(s)

	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			start := space + 1
			end := i - 1

			for start < end {
				bytes[start], bytes[end] = bytes[end], bytes[start]
				start++
				end--
			}

			space = i
		}
	}

	return string(bytes)
}
