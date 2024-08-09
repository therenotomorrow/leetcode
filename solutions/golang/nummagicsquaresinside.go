package golang

import "slices"

func numMagicSquaresInside(grid [][]int) int { //nolint:cyclop
	const (
		gridSize  = 3
		targetSum = 15
	)

	checker := func(startI int, startJ int) bool {
		seen := NewSet[int]()
		currColSum := make([]int, gridSize)
		currRowSum := make([]int, gridSize)
		currDiagSum := 0

		for i := startI; i < gridSize+startI; i++ {
			for j := startJ; j < gridSize+startJ; j++ {
				num := grid[i][j]

				if num < 1 || num > 9 || seen.Contains(num) {
					return false
				}

				seen.Add(num)

				currRowSum[i-startI] += num
				currColSum[j-startJ] += num

				if i-j == startI-startJ {
					currDiagSum += num
				}
			}
		}

		for _, sum := range slices.Concat(currColSum, currRowSum) {
			if sum != targetSum {
				return false
			}
		}

		return currDiagSum == targetSum
	}

	ans := 0

	for i := range len(grid) - 2 {
		for j := range len(grid[i]) - 2 {
			if checker(i, j) {
				ans++
			}
		}
	}

	return ans
}
