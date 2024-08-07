package minimumCost2

import "math"

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	const Alphabet = 26

	minCost := [Alphabet][Alphabet]float64{}
	for i, row := range minCost {
		for j := range row {
			minCost[i][j] = math.MaxInt
		}
	}

	for i := range len(original) {
		sourceS := original[i] - 'a'
		targetS := changed[i] - 'a'

		minCost[sourceS][targetS] = min(minCost[sourceS][targetS], float64(cost[i]))
	}

	for k := range Alphabet {
		for i := range Alphabet {
			for j := range Alphabet {
				minCost[i][j] = min(minCost[i][j], minCost[i][k]+minCost[k][j])
			}
		}
	}

	totalCost := float64(0)

	for i := range len(source) {
		if source[i] == target[i] {
			continue
		}

		sourceS := source[i] - 'a'
		targetS := target[i] - 'a'

		if minCost[sourceS][targetS] >= math.MaxInt {
			return -1
		}

		totalCost += minCost[sourceS][targetS]
	}

	return int64(totalCost)
}
