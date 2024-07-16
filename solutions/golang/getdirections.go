package golang

func getDirections(root *TreeNode, startValue int, destValue int) string { //nolint:funlen,cyclop
	var (
		lca func(node *TreeNode) *TreeNode
		dfs func(node *TreeNode, search int, dir []rune) ([]rune, bool)
	)

	lca = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		if node.Val == startValue || node.Val == destValue {
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

	dfs = func(node *TreeNode, search int, dirs []rune) ([]rune, bool) {
		if node == nil {
			return dirs, false
		}

		if node.Val == search {
			return dirs, true
		}

		dirs, foundLeft := dfs(node.Left, search, dirs)
		if foundLeft {
			return append(dirs, 'L'), foundLeft
		}

		dirs, foundRight := dfs(node.Right, search, dirs)
		if foundRight {
			return append(dirs, 'R'), foundRight
		}

		return dirs, false
	}

	ans := make([]rune, 0)
	common := lca(root)

	dirs, _ := dfs(common, startValue, make([]rune, 0))
	for range dirs {
		ans = append(ans, 'U')
	}

	dirs, _ = dfs(common, destValue, make([]rune, 0))
	for i := len(dirs) - 1; i >= 0; i-- {
		ans = append(ans, dirs[i])
	}

	return string(ans)
}
