package findCheapestPrice

import "math"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[src] = 0

	for i := 0; i <= k; i++ {
		temp := make([]int, n)

		copy(temp, dist)

		for _, flight := range flights {
			if dist[flight[0]] != math.MaxInt {
				temp[flight[1]] = min(temp[flight[1]], dist[flight[0]]+flight[2])
			}
		}

		dist = temp
	}

	if dist[dst] == math.MaxInt {
		return -1
	}

	return dist[dst]
}
