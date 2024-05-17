package golang

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	var dfs func(curr *TreeNode) *TreeNode

	dfs = func(curr *TreeNode) *TreeNode {
		if curr == nil {
			return nil
		}

		curr.Left = dfs(curr.Left)
		curr.Right = dfs(curr.Right)

		if curr.Left == nil && curr.Right == nil && curr.Val == target {
			return nil
		}

		return curr
	}

	return dfs(root)
}
