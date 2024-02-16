package solutions

import "github.com/therenotomorrow/leetcode/internal/structs"

func sumOfLeftLeaves(root *structs.TreeNode) int {
	var traversal func(root *structs.TreeNode, isLeft bool) int

	traversal = func(root *structs.TreeNode, isLeft bool) int {
		if root == nil {
			return 0
		}

		if root.Left == nil && root.Right == nil && isLeft {
			return root.Val
		}

		return traversal(root.Left, true) + traversal(root.Right, false)
	}

	return traversal(root, false)
}
