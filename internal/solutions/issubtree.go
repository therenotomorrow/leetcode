package solutions

import "github.com/therenotomorrow/leetcode/internal/structs"

func isSubtree(root *structs.TreeNode, subRoot *structs.TreeNode) bool {
	if root == nil {
		return false
	}

	return isSameSubtree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameSubtree(root *structs.TreeNode, subRoot *structs.TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == nil && subRoot == nil
	}

	return root.Val == subRoot.Val && isSameSubtree(root.Left, subRoot.Left) && isSameSubtree(root.Right, subRoot.Right)
}
