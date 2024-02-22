package golang

func maxLength(arr []string) int {
	var dynamic func(curr int, str string) int

	dynamic = func(currIdx int, currStr string) int {
		set := NewSet[rune]()
		for _, r := range currStr {
			set.Add(r)
		}

		maxLen := set.Size()

		if maxLen != len(currStr) {
			return 0
		}

		for i := currIdx; i < len(arr); i++ {
			maxLen = Max(maxLen, dynamic(i+1, currStr+arr[i]))
		}

		return maxLen
	}

	return dynamic(0, "")
}
