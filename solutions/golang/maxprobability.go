package golang

func maxProbability(n int, edges [][]int, probs []float64, startNode int, endNode int) float64 {
	// Bellman-Ford Algorithm used
	maxProbs := make([]float64, n)
	maxProbs[startNode] = 1.0 // in start node we successfully could be

	for range n {
		overCycle := true

		for i, edge := range edges {
			prob := probs[i]
			from, dest := edge[0], edge[1]

			switch {
			case maxProbs[from]*prob > maxProbs[dest]:
				maxProbs[dest] = maxProbs[from] * prob
				overCycle = false
			case maxProbs[dest]*prob > maxProbs[from]:
				maxProbs[from] = maxProbs[dest] * prob
				overCycle = false
			}
		}

		if overCycle {
			break
		}
	}

	return maxProbs[endNode]
}
