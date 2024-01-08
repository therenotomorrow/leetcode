package solutions

import "github.com/therenotomorrow/leetcode/pkg/mathfunc"

type BinaryMatrix interface {
	Get(row int, col int) int
	Dimensions() []int
}

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	dims := binaryMatrix.Dimensions()
	leftMost := dims[1]

	for row := 0; row < dims[0]; row++ {
		left := 0
		right := dims[1] - 1

		for left <= right {
			mid := (right + left) / 2

			if binaryMatrix.Get(row, mid) == 0 {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		leftMost = mathfunc.Min(leftMost, left)
	}

	if leftMost == dims[1] {
		return -1
	}

	return leftMost
}
