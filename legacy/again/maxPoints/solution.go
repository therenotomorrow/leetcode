package maxPoints

func maxPoints(points [][]int) int64 {
	rows := len(points)
	cols := len(points[0])

	previousRow := make([]int, cols)

	for col := range cols {
		previousRow[col] = points[0][col]
	}

	for row := range rows - 1 {
		leftMax := make([]int, cols)
		rightMax := make([]int, cols)
		currentRow := make([]int, cols)

		// Calculate left-to-right maximum
		leftMax[0] = previousRow[0]
		for col := 1; col < cols; col++ {
			leftMax[col] = max(leftMax[col-1]-1, previousRow[col])
		}

		// Calculate right-to-left maximum
		rightMax[cols-1] = previousRow[cols-1]
		for col := cols - 2; col >= 0; col-- {
			rightMax[col] = max(rightMax[col+1]-1, previousRow[col])
		}

		// Calculate the current row's maximum points
		for col := 0; col < cols; col++ {
			currentRow[col] = points[row+1][col] + max(leftMax[col], rightMax[col])
		}

		// Update previousRow for the next iteration
		previousRow = currentRow
	}

	// Find the maximum value in the last processed row
	maxPoints := 0
	for col := range cols {
		maxPoints = max(maxPoints, previousRow[col])
	}

	return int64(maxPoints)
}
