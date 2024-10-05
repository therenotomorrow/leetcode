package golang

func checkInclusion(tmpl string, text string) bool {
	if len(tmpl) > len(text) {
		return false
	}

	var (
		cnt1   = make([]int, Alphabet)
		cnt2   = make([]int, Alphabet)
		equals = func(arr1 []int, arr2 []int) bool {
			ans := true

			for i := range Alphabet {
				if arr1[i] != arr2[i] {
					ans = false

					break
				}
			}

			return ans
		}
	)

	for i := range tmpl {
		cnt1[tmpl[i]-'a']++
		cnt2[text[i]-'a']++
	}

	for i := range len(text) - len(tmpl) {
		if equals(cnt1, cnt2) {
			return true
		}

		cnt2[text[i+len(tmpl)]-'a']++
		cnt2[text[i]-'a']--
	}

	return equals(cnt1, cnt2)
}
