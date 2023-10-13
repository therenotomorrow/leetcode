package largestValues

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	result := make([]int, 0)

	if root == nil {
		return result
	}

	var (
		curr  *TreeNode
		queue []*TreeNode
	)

	queue = append(queue, root)

	for len(queue) > 0 {
		mx := math.MinInt
		nm := len(queue)

		for i := 0; i < nm; i++ {
			curr = queue[0]
			queue = queue[1:]

			if curr.Val > mx {
				mx = curr.Val
			}

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		result = append(result, mx)
	}

	return result
}
