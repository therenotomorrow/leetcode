package solutions

import "github.com/therenotomorrow/leetcode/internal/structs"

func invertTree(root *structs.TreeNode) *structs.TreeNode {
	if root == nil {
		return nil
	}

	invertTree(root.Left)
	invertTree(root.Right)

	root.Left, root.Right = root.Right, root.Left

	return root
}
