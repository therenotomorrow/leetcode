package isValidBST

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTConstraint(root, math.MinInt, math.MaxInt)
}

func isValidBSTConstraint(node *TreeNode, min int, max int) bool {
	if node == nil {
		return true
	}

	if !(min < node.Val && node.Val < max) {
		return false
	}

	left := isValidBSTConstraint(node.Left, min, node.Val)
	right := isValidBSTConstraint(node.Right, node.Val, max)

	return left && right
}
