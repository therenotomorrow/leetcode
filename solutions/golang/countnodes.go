package golang

func countNodes(root *TreeNode) int {
	// took from inordertraversal.go: https://leetcode.com/problems/binary-tree-inorder-traversal/description/
	return len(inorderTraversal(root))
}
