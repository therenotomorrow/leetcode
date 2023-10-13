package maxArea

func area(a int, b int) int {
	return a * b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(heights []int) int {
	res := 0

	for i, j := 0, len(heights)-1; i < j; {
		height := min(heights[i], heights[j])

		res = max(area(height, j-i), res)

		if heights[i] < heights[j] {
			i++
		} else {
			j--
		}
	}

	return res
}
