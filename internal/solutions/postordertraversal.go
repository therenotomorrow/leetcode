package solutions

import (
	"github.com/therenotomorrow/leetcode/internal/structs"
	"github.com/therenotomorrow/leetcode/pkg/datastruct"
)

func postorderTraversal(root *structs.TreeNode) []int {
	order := make([]int, 0)
	stack := datastruct.NewStack[*structs.TreeNode]()

	if root == nil {
		return order
	}

	stack.Push(root)

	for !stack.IsEmpty() {
		root, _ = stack.Pop()

		if root == nil {
			continue
		}

		order = append(order, root.Val)

		stack.Push(root.Left)
		stack.Push(root.Right)
	}

	for i, j := 0, len(order)-1; i < j; {
		order[i], order[j] = order[j], order[i]
		i++
		j--
	}

	return order
}
