package golang

func modifyString(s string) string {
	ans := []rune(s)

	for i, runa := range s {
		if runa != '?' {
			continue
		}

		for replacement := 'a'; replacement <= 'z'; replacement++ {
			var left rune
			if i > 0 {
				left = ans[i-1]
			}

			var right rune
			if i < len(ans)-1 {
				right = ans[i+1]
			}

			if left != replacement && replacement != right {
				ans[i] = replacement

				break
			}
		}
	}

	return string(ans)
}
