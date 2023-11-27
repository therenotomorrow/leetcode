package solutions

import (
	"github.com/therenotomorrow/leetcode/internal/structs"
	"github.com/therenotomorrow/leetcode/pkg/cache"
)

func knightDialer(n int) int {
	var (
		c       = cache.NewCache()
		dynamic func(remain int, currDigit int) (numsCnt int)
	)

	graph := [][]int{
		{4, 6},
		{6, 8},
		{7, 9},
		{4, 8},
		{3, 9, 0},
		{},
		{1, 7, 0},
		{2, 6},
		{1, 3},
		{2, 4},
	}

	dynamic = func(remain int, digit int) (numsCnt int) {
		if remain == 0 {
			return 1
		}

		if val, ok := c.Load(remain, digit); ok {
			return val
		}

		for _, nextDigit := range graph[digit] {
			numsCnt = (numsCnt + dynamic(remain-1, nextDigit)) % structs.MOD
		}

		c.Save(numsCnt, remain, digit)

		return numsCnt
	}

	numsCnt := 0
	for digit := 0; digit < 10; digit++ {
		numsCnt = (numsCnt + dynamic(n-1, digit)) % structs.MOD
	}

	return numsCnt
}
