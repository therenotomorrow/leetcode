package zigzagLevelOrder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		num := len(queue)
		res := make([]int, num)

		for i, j := 0, 0; i < num; i++ {
			node := queue[0]
			queue = queue[1:]

			if level%2 == 0 {
				j = i
			} else {
				j = num - i - 1
			}

			res[j] = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, res)
	}

	return result
}
