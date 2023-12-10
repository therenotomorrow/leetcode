package solutions

import (
	"github.com/therenotomorrow/leetcode/internal/structs"
	"github.com/therenotomorrow/leetcode/pkg/datastruct"
)

func preorderTraversal(root *structs.TreeNode) []int {
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

		stack.Push(root.Right)
		stack.Push(root.Left)
	}

	return order
}
