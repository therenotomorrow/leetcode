package golang

func equalToDescendants(root *TreeNode) int {
	var (
		count = 0
		dfs   func(node *TreeNode) int
	)

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftSum := dfs(node.Left)
		rightSum := dfs(node.Right)

		if leftSum+rightSum == node.Val {
			count++
		}

		return leftSum + rightSum + node.Val
	}

	dfs(root)

	return count
}
