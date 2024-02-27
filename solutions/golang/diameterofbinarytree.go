package golang

func diameterOfBinaryTree(root *TreeNode) int {
	var (
		diameter = 0
		dfs      func(node *TreeNode) int
	)

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftPath := dfs(node.Left)
		rightPath := dfs(node.Right)

		diameter = Max(diameter, leftPath+rightPath)

		return Max(leftPath, rightPath) + 1
	}

	dfs(root)

	return diameter
}
