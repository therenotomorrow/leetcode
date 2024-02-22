package golang

func maxAncestorDiff(root *TreeNode) int {
	return maxAncestorDiffCalc(root, root.Val, root.Val)
}

func maxAncestorDiffCalc(node *TreeNode, min int, max int) int {
	if node == nil {
		return max - min
	}

	min = Min(min, node.Val)
	max = Max(max, node.Val)

	lDiff := maxAncestorDiffCalc(node.Left, min, max)
	rDiff := maxAncestorDiffCalc(node.Right, min, max)

	return Max(lDiff, rDiff)
}
