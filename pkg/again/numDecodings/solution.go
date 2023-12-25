package numDecodings

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	var (
		c       = NewCache()
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

type Cache interface {
	Load(keys ...int) (val int, ok bool)
	Save(val int, keys ...int)
}

type cache struct {
	data map[string]int
}

func NewCache() Cache {
	return &cache{data: make(map[string]int)}
}

func (c *cache) makeKey(parts []int) string {
	return fmt.Sprintf("%v", parts)
}

func (c *cache) Load(keys ...int) (val int, ok bool) {
	val, ok = c.data[c.makeKey(keys)]
	return
}

func (c *cache) Save(val int, keys ...int) {
	c.data[c.makeKey(keys)] = val
}
