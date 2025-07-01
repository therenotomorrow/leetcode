package golang

func checkEqualTree(root *TreeNode) bool {
	var (
		sum = make([]int, 0)
		dfs func(node *TreeNode) int
	)

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		val := dfs(node.Left) + dfs(node.Right) + node.Val
		sum = append(sum, val)

		return val
	}

	totalSum := dfs(root)
	if totalSum%Half != 0 {
		return false
	}

	halfSum := totalSum / Half
	sum = sum[:len(sum)-1] // remove total - it doesn't count

	for _, num := range sum {
		if num == halfSum {
			return true
		}
	}

	return false
}
