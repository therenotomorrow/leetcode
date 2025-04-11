package golang

type FindElements struct {
	data Set[int]
}

func FindElementsConstructor(root *TreeNode) FindElements {
	var (
		set = NewSet[int]()
		dfs func(node *TreeNode) Set[int]
	)

	dfs = func(node *TreeNode) Set[int] {
		if node == nil {
			return set
		}

		set.Add(node.Val)

		if node.Left != nil {
			node.Left.Val = Double*node.Val + 1
			dfs(node.Left)
		}

		if node.Right != nil {
			node.Right.Val = Double*node.Val + Double
			dfs(node.Right)
		}

		return set
	}

	// start with the zero
	root.Val = 0

	return FindElements{data: dfs(root)}
}

func (fe *FindElements) Find(target int) bool {
	return fe.data.Contains(target)
}
