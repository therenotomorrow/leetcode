package golang

func maxAncestorDiff(root *TreeNode) int {
	return maxAncestorDiffCalc(root, root.Val, root.Val)
}

func maxAncestorDiffCalc(node *TreeNode, mini int, maxi int) int {
	if node == nil {
		return maxi - mini
	}

	mini = Min(mini, node.Val)
	maxi = Max(maxi, node.Val)

	lDiff := maxAncestorDiffCalc(node.Left, mini, maxi)
	rDiff := maxAncestorDiffCalc(node.Right, mini, maxi)

	return Max(lDiff, rDiff)
}
