package golang

func findDistance(root *TreeNode, valP int, valQ int) int {
	var (
		lca func(node *TreeNode) *TreeNode
		dfs func(node *TreeNode, search int, depth int) int
	)

	lca = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		if node.Val == valP || node.Val == valQ {
			return node
		}

		left := lca(node.Left)
		right := lca(node.Right)

		if left != nil && right != nil {
			return node
		}

		if left != nil {
			return left
		}

		return right
	}

	dfs = func(node *TreeNode, search int, depth int) int {
		if node == nil {
			return -1
		}

		if node.Val == search {
			return depth
		}

		left := dfs(node.Left, search, depth+1)
		right := dfs(node.Right, search, depth+1)

		return Max(left, right)
	}

	common := lca(root)

	return dfs(common, valP, 0) + dfs(common, valQ, 0)
}
