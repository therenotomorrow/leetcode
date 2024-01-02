package numOfArrays

import "fmt"

const MOD = 1e9 + 7

type Cache struct {
	data map[string]int
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]int)}
}

func (c *Cache) key(i int, j int, k int) string {
	return fmt.Sprintf("%d,%d,%d", i, j, k)
}

func (c *Cache) Save(i int, j int, k int, res int) int {
	c.data[c.key(i, j, k)] = res
	return res
}

func (c *Cache) Load(i int, j int, k int) (int, bool) {
	val, ok := c.data[c.key(i, j, k)]
	return val, ok
}

func numOfArrays(n int, m int, k int) int {
	cache := NewCache()

	var dynamic func(int, int, int) int

	dynamic = func(i int, maxSoFar int, remain int) int {
		if i == n {
			if remain == 0 {
				return 1
			} else {
				return 0
			}
		}

		if remain < 0 {
			return 0
		}

		if val, ok := cache.Load(i, maxSoFar, remain); ok {
			return val
		}

		ans := 0
		for num := 1; num <= maxSoFar; num++ {
			ans = (ans + dynamic(i+1, maxSoFar, remain)) % MOD
		}

		for num := maxSoFar + 1; num <= m; num++ {
			ans = (ans + dynamic(i+1, num, remain-1)) % MOD
		}

		return cache.Save(i, maxSoFar, remain, ans)
	}

	return dynamic(0, 0, k)
}
