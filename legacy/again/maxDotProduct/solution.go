package maxDotProduct

import "fmt"

func max(nums ...int) int {
	m := nums[0]

	for _, num := range nums[1:] {
		if num > m {
			m = num
		}
	}

	return m
}

func min(nums ...int) int {
	m := nums[0]

	for _, num := range nums[1:] {
		if num < m {
			m = num
		}
	}

	return m
}

type Cache struct {
	data map[string]int
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]int)}
}

func (c *Cache) key(i int, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}

func (c *Cache) Save(i int, j int, res int) {
	c.data[c.key(i, j)] = res
}

func (c *Cache) Load(i int, j int) (int, bool) {
	val, ok := c.data[c.key(i, j)]
	return val, ok
}

func maxDotProduct(nums1 []int, nums2 []int) int {
	cache := NewCache()

	if max(nums1...) < 0 && min(nums2...) > 0 {
		return max(nums1...) * min(nums2...)
	}

	if min(nums1...) > 0 && max(nums2...) < 0 {
		return min(nums1...) * max(nums2...)
	}

	var dynamic func(int, int) int

	dynamic = func(i int, j int) int {
		if val, ok := cache.Load(i, j); ok {
			return val
		}

		if i == len(nums1) || j == len(nums2) {
			return 0
		}

		use := nums1[i]*nums2[j] + dynamic(i+1, j+1)

		ans := max(use, dynamic(i+1, j), dynamic(i, j+1))

		cache.Save(i, j, ans)

		return ans
	}

	return dynamic(0, 0)
}
