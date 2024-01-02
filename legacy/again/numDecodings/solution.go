package numDecodings

import (
	"strconv"

	"github.com/therenotomorrow/leetcode/pkg/cache"
)

func numDecodings(s string) int {
	var (
		c       = cache.NewCache()
		dynamic func(idx int) (cnt int)
	)

	dynamic = func(idx int) (cnt int) {
		if val, ok := c.Load(idx); ok {
			return val
		}

		cnt = -1

		switch {
		case idx == len(s):
			cnt = 1
		case s[idx] == '0':
			cnt = 0
		case idx == len(s)-1:
			cnt = 1
		}

		if cnt != -1 {
			return cnt
		}

		cnt = dynamic(idx + 1)

		if val, _ := strconv.Atoi(s[idx : idx+2]); val <= 26 {
			cnt += dynamic(idx + 2)
		}

		c.Save(cnt, idx)

		return cnt
	}

	return dynamic(0)
}
