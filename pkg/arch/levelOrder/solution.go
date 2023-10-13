package levelOrder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ans := make([][]int, 0)

	for queue := []*TreeNode{root}; len(queue) > 0; {
		levelAns := make([]int, len(queue))

		for i := 0; i < len(levelAns); i++ {
			node := queue[0]

			levelAns[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			queue = queue[1:]
		}

		ans = append(ans, levelAns)
	}

	return ans
}
