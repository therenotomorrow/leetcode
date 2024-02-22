package golang

func treeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := treeHeight(root.Left)
	right := treeHeight(root.Right)

	return 1 + Max(left, right)
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lh := treeHeight(root.Left)
	rh := treeHeight(root.Right)

	if Abs(rh-lh) < 2 && isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}

	return false
}
