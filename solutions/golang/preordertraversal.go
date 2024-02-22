package golang

func preorderTraversal(root *TreeNode) []int {
	order := make([]int, 0)
	stack := NewStack[*TreeNode]()

	if root == nil {
		return order
	}

	stack.Push(root)

	for !stack.IsEmpty() {
		root, _ = stack.Pop()

		if root == nil {
			continue
		}

		order = append(order, root.Val)

		stack.Push(root.Right)
		stack.Push(root.Left)
	}

	return order
}
