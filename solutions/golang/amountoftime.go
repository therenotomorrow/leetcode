package golang

func amountOfTime(root *TreeNode, start int) int {
	var time int

	traversal(root, start, &time)

	return time
}

func traversal(root *TreeNode, start int, time *int) int {
	if root == nil {
		return 0
	}

	left := traversal(root.Left, start, time)
	right := traversal(root.Right, start, time)

	switch {
	case root.Val == start:
		*time = Max(left, right)

		return -1 // signal outside that we found infected node

	case left > -1 && right > -1:
		return Max(left, right) + 1

	default:
		*time = Max(*time, Abs(left)+Abs(right))

		return Min(left, right) - 1 // how far infected node from root
	}
}
