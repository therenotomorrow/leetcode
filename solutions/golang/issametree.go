package golang

func isSameTree(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	return isSameTree(left.Left, right.Left) && isSameTree(left.Right, right.Right)
}
