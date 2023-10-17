package cache

import "fmt"

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
