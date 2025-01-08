package golang

func equalDigitFrequency(text string) int {
	set := NewSet[string]()

	counter := func(str string) bool {
		cnt := make([]int, Digits)
		for _, ch := range str {
			cnt[ch-'0']++
		}

		uniq := Max(cnt...)
		for _, times := range cnt {
			if times != 0 && uniq != times {
				return false
			}
		}

		return true
	}

	for i := range text {
		for j := i; j < len(text); j++ {
			str := text[i : j+1]

			if counter(str) {
				set.Add(str)
			}
		}
	}

	return set.Size()
}
