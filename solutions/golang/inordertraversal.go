package golang

func inorderTraversal(root *TreeNode) []int {
	order := make([]int, 0)
	stack := NewStack[*TreeNode]()

	if root == nil {
		return order
	}

	for root != nil || !stack.IsEmpty() {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}

		root, _ = stack.Pop()
		order = append(order, root.Val)

		root = root.Right
	}

	return order
}
