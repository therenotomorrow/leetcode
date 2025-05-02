package golang

import "fmt"

type Cacheable interface {
	Load(keys ...int) (int, bool)
	Save(val int, keys ...int)
}

type Cache struct {
	data map[string]int
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]int)}
}

func (c *Cache) Load(keys ...int) (int, bool) {
	val, ok := c.data[c.makeKey(keys)]

	return val, ok
}

func (c *Cache) Save(val int, keys ...int) {
	c.data[c.makeKey(keys)] = val
}

func (c *Cache) makeKey(parts []int) string {
	return fmt.Sprintf("%v", parts)
}
