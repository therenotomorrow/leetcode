package modifiedGraphEdges

const INF = int(2e9)

func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
	// Step 1: Compute the initial shortest distance from source to destination
	currentShortestDistance := runDijkstra(
		edges,
		n,
		source,
		destination)

	// If the current shortest distance is less than the target, return an empty result
	if currentShortestDistance < target {
		return [][]int{}
	}

	matchesTarget := currentShortestDistance == target

	// Step 2: Iterate through each edge to adjust its weight if necessary
	for _, edge := range edges {
		// Skip edges that already have a positive weight
		if edge[2] > 0 {
			continue
		}

		// Set edge weight to a large value if current distance matches target, else set to 1
		if matchesTarget {
			edge[2] = INF
		} else {
			edge[2] = 1
		}

		if !matchesTarget {
			// Compute the new shortest distance with the updated edge weight
			newDistance := runDijkstra(edges, n, source, destination)
			// If the new distance is within the target range, update edge weight to match target
			if newDistance <= target {
				matchesTarget = true
				edge[2] += target - newDistance
			}
		}
	}

	if matchesTarget {
		return edges
	}

	// Return modified edges if the target distance is achieved, otherwise return an empty result
	return [][]int{}

}

func runDijkstra(edges [][]int, n int, source int, destination int) int {
	// Step 1: Initialize adjacency matrix and distance arrays
	adjMatrix := make([][]int, n)
	for i := range adjMatrix {
		adjMatrix[i] = make([]int, n)
		for j := range adjMatrix[i] {
			adjMatrix[i][j] = INF
		}
	}

	minDistance := make([]int, n)
	for i := range minDistance {
		minDistance[i] = INF
	}

	visited := make([]bool, n)

	// Set the distance to the source node as 0
	minDistance[source] = 0

	// Step 2: Fill the adjacency matrix with edge weights
	for _, edge := range edges {
		if edge[2] != -1 {
			adjMatrix[edge[0]][edge[1]] = edge[2]
			adjMatrix[edge[1]][edge[0]] = edge[2]
		}
	}

	// Step 3: Perform Dijkstra's algorithm
	for range n {
		// Find the nearest unvisited node
		nearestUnvisitedNode := -1
		for j := range n {
			if !visited[j] && (nearestUnvisitedNode == -1 || minDistance[j] < minDistance[nearestUnvisitedNode]) {
				nearestUnvisitedNode = j
			}
		}
		// Mark the nearest node as visited
		visited[nearestUnvisitedNode] = true

		// Update the minimum distance for each adjacent node
		for v := range n {
			minDistance[v] = min(minDistance[v], minDistance[nearestUnvisitedNode]+adjMatrix[nearestUnvisitedNode][v])
		}
	}

	// Return the shortest distance to the destination node
	return minDistance[destination]
}
