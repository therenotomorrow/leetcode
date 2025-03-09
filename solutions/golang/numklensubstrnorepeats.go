package golang

func numKLenSubstrNoRepeats(text string, k int) int {
	if len(text) < k {
		return 0
	}

	left := 0
	ans := 0
	cnt := make(map[byte]int)

	for right := range text {
		cnt[text[right]]++

		if right-left+1 < k {
			continue
		}

		if len(cnt) == k {
			ans++
		}

		cnt[text[left]]--
		if cnt[text[left]] == 0 {
			delete(cnt, text[left])
		}

		left++
	}

	return ans
}
