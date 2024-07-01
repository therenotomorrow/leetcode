package golang

func isValidPalindrome(str string, remove int) bool {
	var (
		cache   = NewCache()
		dynamic func(left int, right int) int
	)

	dynamic = func(left int, right int) int {
		if val, ok := cache.Load(left, right); ok {
			return val
		}

		switch {
		case left == right:
			return 0
		case right-left == 1:
			if str[left] != str[right] {
				return 1
			} else {
				return 0
			}
		case str[left] == str[right]:
			return dynamic(left+1, right-1)
		}

		removesCnt := 1 + Min(dynamic(left+1, right), dynamic(left, right-1))

		cache.Save(removesCnt, left, right)

		return removesCnt
	}

	return dynamic(0, len(str)-1) <= remove
}
