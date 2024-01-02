package findDiagonalOrder

func reverseArr(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}

func findDiagonalOrder(mat [][]int) []int {
	var (
		numRows     = len(mat)
		numCols     = len(mat[0])
		fromBotToUp = true
		arr         = make([]int, 0, numRows*numCols)
	)

	//  1   2   3  4  5
	//  6   7   8  9  -
	// 11  12  13  -  -
	// 16  17   -  -  -
	for i := 0; i < numCols; i++ {
		tmp := make([]int, 0)

		for j := 0; j <= i && j < numRows; j++ {
			tmp = append(tmp, mat[j][i-j])
		}

		if fromBotToUp {
			reverseArr(tmp)
		}

		arr = append(arr, tmp...)
		fromBotToUp = !fromBotToUp
	}

	// draw the last triangle
	for i := 1; i < numRows; i++ {
		tmp := make([]int, 0)

		for j := numCols - 1; j >= i+numCols-numRows && j >= 0; j-- {
			tmp = append(tmp, mat[numCols-j-1+i][j])
		}

		if fromBotToUp {
			reverseArr(tmp)
		}

		arr = append(arr, tmp...)
		fromBotToUp = !fromBotToUp
	}

	return arr
}
