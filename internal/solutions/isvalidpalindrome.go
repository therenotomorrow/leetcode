package solutions

import (
	"github.com/therenotomorrow/leetcode/pkg/cache"
	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
)

func isValidPalindrome(s string, k int) bool {
	var (
		c       = cache.NewCache()
		dynamic func(left int, right int) (removesCnt int)
	)

	dynamic = func(left int, right int) (removesCnt int) {
		if val, ok := c.Load(left, right); ok {
			return val
		}

		switch {
		case left == right:
			return 0
		case right-left == 1:
			if s[left] != s[right] {
				return 1
			} else {
				return 0
			}
		case s[left] == s[right]:
			return dynamic(left+1, right-1)
		}

		removesCnt = 1 + mathfunc.Min(dynamic(left+1, right), dynamic(left, right-1))

		c.Save(removesCnt, left, right)

		return removesCnt
	}

	return dynamic(0, len(s)-1) <= k
}
