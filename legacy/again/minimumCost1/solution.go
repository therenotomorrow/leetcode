package minimumCost1

import (
	"container/heap"
	"math"
)

type Edge struct {
	city int
	toll int
}

type Item struct {
	cost      int
	city      int
	discounts int
}

type PriQueue []Item

func (pq *PriQueue) Len() int { return len(*pq) }

func (pq *PriQueue) Less(i, j int) bool {
	return (*pq)[i].cost < (*pq)[j].cost
}

func (pq *PriQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriQueue) Push(x interface{}) {
	item := x.(Item)
	*pq = append(*pq, item)
}

func (pq *PriQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func minimumCost(n int, highways [][]int, discounts int) int {
	graph := make([][]Edge, n)
	for _, highway := range highways {
		graph[highway[0]] = append(graph[highway[0]], Edge{highway[1], highway[2]})
		graph[highway[1]] = append(graph[highway[1]], Edge{highway[0], highway[2]})
	}

	pq := &PriQueue{}

	heap.Init(pq)
	heap.Push(pq, Item{0, 0, 0})

	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, discounts+1)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt
		}
	}

	dist[0][0] = 0

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, discounts+1)
	}

	for pq.Len() > 0 {
		current := heap.Pop(pq).(Item)

		if visited[current.city][current.discounts] {
			continue
		}

		visited[current.city][current.discounts] = true

		for _, neighbor := range graph[current.city] {
			neighborCity, toll := neighbor.city, neighbor.toll

			if current.cost+toll < dist[neighborCity][current.discounts] {
				dist[neighborCity][current.discounts] = current.cost + toll
				heap.Push(pq, Item{dist[neighborCity][current.discounts], neighborCity, current.discounts})
			}

			if current.discounts < discounts {
				newCostWithDiscount := current.cost + toll/2
				if newCostWithDiscount < dist[neighborCity][current.discounts+1] {
					dist[neighborCity][current.discounts+1] = newCostWithDiscount
					heap.Push(pq, Item{newCostWithDiscount, neighborCity, current.discounts + 1})
				}
			}
		}
	}

	minCost := math.MaxInt
	for i := 0; i <= discounts; i++ {
		if dist[n-1][i] < minCost {
			minCost = dist[n-1][i]
		}
	}

	if minCost == math.MaxInt {
		minCost = -1
	}

	return minCost
}
