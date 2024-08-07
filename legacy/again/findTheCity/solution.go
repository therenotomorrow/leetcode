package findTheCity

import (
	"container/heap"
	"math"
)

type Edge struct {
	to, weight int
}

type PriorityQueueItem struct {
	distance, city int
}

type PriorityQueue []PriorityQueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(PriorityQueueItem)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	adjacencyList := make([][]Edge, n)
	shortestPathMatrix := make([][]float64, n)
	for i := range shortestPathMatrix {
		shortestPathMatrix[i] = make([]float64, n)
		for j := range shortestPathMatrix[i] {
			shortestPathMatrix[i][j] = math.Inf(1)
		}
		shortestPathMatrix[i][i] = 0
	}

	for _, edge := range edges {
		start, end, weight := edge[0], edge[1], edge[2]
		adjacencyList[start] = append(adjacencyList[start], Edge{end, weight})
		adjacencyList[end] = append(adjacencyList[end], Edge{start, weight})
	}

	for i := 0; i < n; i++ {
		dijkstra(n, adjacencyList, shortestPathMatrix[i], i)
	}

	return getCityWithFewestReachable(n, shortestPathMatrix, distanceThreshold)
}

func dijkstra(n int, adjacencyList [][]Edge, shortestPathDistances []float64, source int) {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, PriorityQueueItem{0, source})

	for i := range shortestPathDistances {
		shortestPathDistances[i] = math.Inf(1)
	}
	shortestPathDistances[source] = 0

	for pq.Len() > 0 {
		current := heap.Pop(pq).(PriorityQueueItem)
		currentDistance, currentCity := current.distance, current.city

		if float64(currentDistance) > shortestPathDistances[currentCity] {
			continue
		}

		for _, neighbor := range adjacencyList[currentCity] {
			newDistance := float64(currentDistance) + float64(neighbor.weight)
			if newDistance < shortestPathDistances[neighbor.to] {
				shortestPathDistances[neighbor.to] = newDistance
				heap.Push(pq, PriorityQueueItem{int(newDistance), neighbor.to})
			}
		}
	}
}

func getCityWithFewestReachable(n int, shortestPathMatrix [][]float64, distanceThreshold int) int {
	cityWithFewestReachable := -1
	fewestReachableCount := n

	for i := 0; i < n; i++ {
		reachableCount := 0
		for j := 0; j < n; j++ {
			if i != j && shortestPathMatrix[i][j] <= float64(distanceThreshold) {
				reachableCount++
			}
		}
		if reachableCount <= fewestReachableCount {
			fewestReachableCount = reachableCount
			cityWithFewestReachable = i
		}
	}

	return cityWithFewestReachable
}
