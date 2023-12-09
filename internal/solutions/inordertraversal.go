package solutions

import (
	"github.com/therenotomorrow/leetcode/internal/structs"
	"github.com/therenotomorrow/leetcode/pkg/datastruct"
)

func inorderTraversal(root *structs.TreeNode) []int {
	order := make([]int, 0)
	stack := datastruct.NewStack[*structs.TreeNode]()

	if root == nil {
		return order
	}

	for {
		if root != nil {
			stack.Push(root)
			root = root.Left
			continue
		}

		if stack.IsEmpty() {
			break
		}

		root, _ = stack.Pop()
		order = append(order, root.Val)

		root = root.Right
	}

	return order
}
