package solutions

import "github.com/therenotomorrow/leetcode/internal/structs"

func countNodes(root *structs.TreeNode) int {
	// took from inordertraversal.go: https://leetcode.com/problems/binary-tree-inorder-traversal/description/
	return len(inorderTraversal(root))
}
