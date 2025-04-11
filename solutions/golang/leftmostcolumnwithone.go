package golang

type BinaryMatrix interface {
	Get(row int, col int) int
	Dimensions() []int
}

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	dims := binaryMatrix.Dimensions()
	leftMost := dims[1]

	for row := range dims[0] {
		left := 0
		right := dims[1] - 1

		for left <= right {
			mid := (right + left) / Half

			if binaryMatrix.Get(row, mid) == 0 {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		leftMost = Min(leftMost, left)
	}

	if leftMost == dims[1] {
		return -1
	}

	return leftMost
}
