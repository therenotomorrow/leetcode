package golang

func kthDistinct(arr []string, k int) string {
	cnt := make(map[string]int)

	for _, str := range arr {
		cnt[str]++
	}

	for _, str := range arr {
		if cnt[str] == 1 {
			k--
		}

		if k == 0 {
			return str
		}
	}

	return ""
}
