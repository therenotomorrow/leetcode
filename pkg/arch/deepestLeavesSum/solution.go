package deepestLeavesSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	var (
		levelQueue, nextQueue []*TreeNode
		result, levelSum      int
	)

	nextQueue = append(nextQueue, root)

	for len(nextQueue) > 0 {
		levelQueue = nextQueue
		nextQueue = make([]*TreeNode, 0)

		for levelSum = 0; len(levelQueue) > 0; {
			node := levelQueue[0]
			levelQueue = levelQueue[1:]

			levelSum += node.Val

			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}

			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}

		result = levelSum
	}

	return result
}
