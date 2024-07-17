package golang

func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
	var (
		forest = make([]*TreeNode, 0)
		cnt    = make(map[int]bool)
		dfs    func(node *TreeNode) *TreeNode
	)

	for _, val := range toDelete {
		cnt[val] = true
	}

	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return node
		}

		node.Left = dfs(node.Left)
		node.Right = dfs(node.Right)

		if cnt[node.Val] {
			if node.Left != nil {
				forest = append(forest, node.Left)
			}

			if node.Right != nil {
				forest = append(forest, node.Right)
			}

			return nil
		}

		return node
	}

	root = dfs(root)

	if root != nil {
		forest = append(forest, root)
	}

	return forest
}
