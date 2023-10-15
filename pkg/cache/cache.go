package cache

import "fmt"

type Cache map[string]int

func New() *Cache {
	c := Cache(make(map[string]int))
	return &c
}

func (c *Cache) makeKey(parts []int) string {
	return fmt.Sprintf("%v", parts)
}

func (c *Cache) Load(keys ...int) (int, bool) {
	val, ok := (*c)[c.makeKey(keys)]
	return val, ok
}

func (c *Cache) Save(res int, keys ...int) int {
	(*c)[c.makeKey(keys)] = res
	return res
}
