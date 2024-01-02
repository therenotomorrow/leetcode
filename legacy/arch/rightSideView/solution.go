package rightSideView

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
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
		numNodes := len(queue)

		for i := 0; i < numNodes; i++ {
			curr = queue[0]
			queue = queue[1:]

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		result = append(result, curr.Val)
	}

	return result
}
