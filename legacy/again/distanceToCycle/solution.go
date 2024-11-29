package main

func distanceToCycle(n int, edges [][]int) []int {
	isInCycle := make([]bool, n)
	visited := make([]bool, n)
	parent := make([]int, n)
	distances := make([]int, n)
	adjacencyList := make([][]int, n)

	// Initialize adjacency list
	for i := range adjacencyList {
		adjacencyList[i] = make([]int, 0)
	}

	// Build the graph
	for _, edge := range edges {
		adjacencyList[edge[0]] = append(adjacencyList[edge[0]], edge[1])
		adjacencyList[edge[1]] = append(adjacencyList[edge[1]], edge[0])
	}

	// Detect cycle nodes
	for i := 0; i < n; i++ {
		if !visited[i] {
			if detectCycleNodes(i, -1, adjacencyList, isInCycle, visited, parent) {
				break
			}
		}
	}

	// Reset visited for distance calculation
	visited = make([]bool, n)

	// Calculate distances
	queue := []int{}
	for i := 0; i < n; i++ {
		if isInCycle[i] {
			distances[i] = 0
			queue = append(queue, i)
			visited[i] = true
		}
	}

	// BFS to calculate distances from cycle
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		for _, neighbor := range adjacencyList[currentNode] {
			if !visited[neighbor] {
				visited[neighbor] = true
				distances[neighbor] = distances[currentNode] + 1
				queue = append(queue, neighbor)
			}
		}
	}

	return distances
}

func detectCycleNodes(currentNode, parentNode int, adjacencyList [][]int, isInCycle, visited []bool, parent []int) bool {
	visited[currentNode] = true

	for _, neighbor := range adjacencyList[currentNode] {
		if !visited[neighbor] {
			parent[neighbor] = currentNode
			if detectCycleNodes(neighbor, currentNode, adjacencyList, isInCycle, visited, parent) {
				return true
			}
		} else if neighbor != parentNode {
			// Cycle detected
			isInCycle[neighbor] = true
			tempNode := currentNode
			for tempNode != neighbor {
				isInCycle[tempNode] = true
				tempNode = parent[tempNode]
			}
			return true
		}
	}

	return false
}
