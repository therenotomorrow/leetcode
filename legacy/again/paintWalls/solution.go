package paintWalls

type Cache struct {
	data map[[2]int]int
}

func New() *Cache {
	return &Cache{data: make(map[[2]int]int)}
}

func (c *Cache) Save(val int, keys ...int) int {
	c.data[[2]int{keys[0], keys[1]}] = val
	return val
}

func (c *Cache) Load(keys ...int) (int, bool) {
	val, ok := c.data[[2]int{keys[0], keys[1]}]
	return val, ok
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func paintWalls(cost []int, time []int) int {
	var (
		cache   = New()
		dynamic func(int, int) int
		n       = len(time)
	)

	dynamic = func(i int, remainWalls int) int {
		if val, ok := cache.Load(i, remainWalls); ok {
			return val
		}

		if remainWalls <= 0 {
			return 0
		}

		if i == n {
			return 1_000_000_000
		}

		// -1 because paid will paint 1 wall of time[i],
		// and free can that time[i] paint time[i] walls because of 1 time per wall
		paint := cost[i] + dynamic(i+1, remainWalls-1-time[i])
		notPaint := dynamic(i+1, remainWalls)

		return cache.Save(min(paint, notPaint), i, remainWalls)
	}

	return dynamic(0, n)
}
