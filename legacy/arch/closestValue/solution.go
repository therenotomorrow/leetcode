package closestValue

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func calcDiff(a float64, b float64) float64 {
	if a < b { // b always positive because of task description
		return b - a
	} else {
		return a - b
	}
}

func closestValue(root *TreeNode, target float64) int {
	closed := root.Val
	diff := calcDiff(float64(closed), target)

	for root != nil {
		if d := calcDiff(float64(root.Val), target); d < diff {
			closed = root.Val
			diff = d
		} else if d == diff && root.Val < closed {
			closed = root.Val
			diff = d
		}

		if target < float64(root.Val) {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return closed
}
