package golang

func treeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := treeHeight(root.Left)
	right := treeHeight(root.Right)

	return 1 + Max(left, right)
}

func isBalanced1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lh := treeHeight(root.Left)
	rh := treeHeight(root.Right)

	if Abs(rh-lh) < 2 && isBalanced1(root.Left) && isBalanced1(root.Right) {
		return true
	}

	return false
}
