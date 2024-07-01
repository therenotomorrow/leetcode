package Graph

import (
	"container/heap"
	"math"
)

type item struct {
	node int
	cost int
}

type PriorityQueue struct {
	data []item
}

func (piq *PriorityQueue) Len() int {
	return len(piq.data)
}

func (piq *PriorityQueue) Less(i, j int) bool {
	return piq.data[i].cost < piq.data[j].cost
}

func (piq *PriorityQueue) Swap(i, j int) {
	piq.data[i], piq.data[j] = piq.data[j], piq.data[i]
}

func (piq *PriorityQueue) Push(x any) {
	piq.data = append(piq.data, x.(item))
}

func (piq *PriorityQueue) Pop() any {
	old := piq.data
	n := len(old)
	item := old[n-1]
	piq.data = old[:piq.Len()-1]
	return item
}

func intsToEdge(ints []int) (from int, to int, cost int) {
	from = ints[0]
	to = ints[1]
	cost = ints[2]

	return
}

type Graph struct {
	data [][]item
}

func Constructor(n int, edges [][]int) Graph {
	g := Graph{data: make([][]item, n)}

	for i := 0; i < n; i++ {
		g.data[i] = make([]item, 0)
	}

	for _, edge := range edges {
		from, to, cost := intsToEdge(edge)

		g.data[from] = append(g.data[from], item{node: to, cost: cost})
	}

	return g
}

func (g *Graph) AddEdge(edge []int) {
	from, to, cost := intsToEdge(edge)

	g.data[from] = append(g.data[from], item{node: to, cost: cost})
}

func (g *Graph) ShortestPath(node1 int, node2 int) int {
	n := len(g.data)
	piq := new(PriorityQueue)
	costForNode := make([]int, n)

	for i := range costForNode {
		costForNode[i] = math.MaxInt
	}
	costForNode[node1] = 0

	heap.Push(piq, item{node: node1, cost: 0})

	for piq.Len() > 0 {
		top := heap.Pop(piq).(item)
		currCost := top.cost
		currNode := top.node

		if currCost > costForNode[currNode] {
			continue
		}
		if currNode == node2 {
			return currCost
		}

		for _, neighbor := range g.data[currNode] {
			neighborNode := neighbor.node
			neighborCost := neighbor.cost
			newCost := currCost + neighborCost

			if newCost < costForNode[neighborNode] {
				costForNode[neighborNode] = newCost
				heap.Push(piq, item{node: neighborNode, cost: newCost})
			}
		}
	}

	return -1
}
