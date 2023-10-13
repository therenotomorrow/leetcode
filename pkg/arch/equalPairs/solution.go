package equalPairs

import (
	"strconv"
	"strings"
)

func equalPairs(grid [][]int) int {
	n := len(grid)
	hash := make(map[string]int)

	key := make([]string, n)
	for _, raw := range grid {
		for col := 0; col < n; col++ {
			key[col] = strconv.Itoa(raw[col])
		}

		hash[strings.Join(key, ",")]++
	}

	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			key[j] = strconv.Itoa(grid[j][i])
		}

		cnt += hash[strings.Join(key, ",")]
	}

	return cnt
}
