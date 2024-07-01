package golang

func evaluateTree(root *TreeNode) bool {
	const magicTwo = 2

	if root.Left == nil && root.Right == nil {
		return root.Val == 1
	}

	left := evaluateTree(root.Left)
	right := evaluateTree(root.Right)

	if root.Val == magicTwo {
		return left || right
	}

	return left && right
}
