package golang

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	return isSameSubtree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == nil && subRoot == nil
	}

	return root.Val == subRoot.Val && isSameSubtree(root.Left, subRoot.Left) && isSameSubtree(root.Right, subRoot.Right)
}
