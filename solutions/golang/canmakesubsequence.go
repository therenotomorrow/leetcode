package golang

func canMakeSubsequence(str1 string, str2 string) bool {
	if len(str1) < len(str2) {
		return false
	}

	j := 0

	for i := 0; i < len(str1) && j < len(str2); i++ {
		inc := str1[i] + 1

		if inc > 'z' {
			inc = 'a'
		}

		if str1[i] == str2[j] || inc == str2[j] {
			j++
		}
	}

	return j == len(str2)
}
