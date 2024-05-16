package golang

func evaluateTree(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return root.Val == 1
	}

	left := evaluateTree(root.Left)
	right := evaluateTree(root.Right)

	if root.Val == 2 {
		return left || right
	}

	return left && right
}
