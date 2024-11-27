package shortestDistanceAfterQueries

import (
	"container/heap"
	"math"
)

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	type Edge struct {
		to, weight int
	}

	graph := make([][]Edge, n)
	for i := 0; i < n-1; i++ {
		graph[i] = append(graph[i], Edge{i + 1, 1})
	}

	addEdge := func(u, v int) {
		graph[u] = append(graph[u], Edge{v, 1})
	}

	dijkstra := func() int {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = math.MaxInt32
		}
		dist[0] = 0

		h := &MinHeap{}
		heap.Init(h)
		heap.Push(h, [2]int{0, 0}) // {distance, node}

		for h.Len() > 0 {
			cur := heap.Pop(h).([2]int)
			curDist, node := cur[0], cur[1]

			if curDist > dist[node] {
				continue
			}

			for _, edge := range graph[node] {
				newDist := curDist + edge.weight
				if newDist < dist[edge.to] {
					dist[edge.to] = newDist
					heap.Push(h, [2]int{newDist, edge.to})
				}
			}
		}
		if dist[n-1] == math.MaxInt32 {
			return -1
		}
		return dist[n-1]
	}

	result := []int{}
	for _, query := range queries {
		u, v := query[0], query[1]
		addEdge(u, v)
		result = append(result, dijkstra())
	}

	return result
}

type MinHeap [][2]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
