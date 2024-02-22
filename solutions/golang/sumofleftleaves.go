package golang

func sumOfLeftLeaves(root *TreeNode) int {
	var traversal func(root *TreeNode, isLeft bool) int

	traversal = func(root *TreeNode, isLeft bool) int {
		if root == nil {
			return 0
		}

		if root.Left == nil && root.Right == nil && isLeft {
			return root.Val
		}

		return traversal(root.Left, true) + traversal(root.Right, false)
	}

	return traversal(root, false)
}
