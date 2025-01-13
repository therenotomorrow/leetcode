package golang

func minimumLength2(s string) int {
	cnt := make([]int, Alphabet)
	for _, r := range s {
		cnt[r-'a']++
	}

	ans := 0

	for _, times := range cnt {
		switch {
		case times == 0:
			continue
		case times%2 == 0:
			ans += 2
		default:
			ans++
		}
	}

	return ans
}
