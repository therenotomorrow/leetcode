package golang

import "math"

func largestBSTSubtree(root *TreeNode) int {
	var (
		count    func(node *TreeNode) int
		minVal   func(node *TreeNode) int
		maxVal   func(node *TreeNode) int
		validate func(node *TreeNode) bool
	)

	count = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		return 1 + count(node.Left) + count(node.Right)
	}

	minVal = func(node *TreeNode) int {
		if node == nil {
			return math.MaxInt
		}

		return Min(node.Val, minVal(node.Left), minVal(node.Right))
	}

	maxVal = func(node *TreeNode) int {
		if node == nil {
			return math.MinInt
		}

		return Max(node.Val, maxVal(node.Left), maxVal(node.Right))
	}

	validate = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		if lMax := maxVal(node.Left); lMax >= node.Val {
			return false
		}

		if rMin := minVal(node.Right); rMin <= node.Val {
			return false
		}

		return validate(node.Left) && validate(node.Right)
	}

	if root == nil {
		return 0
	}

	if validate(root) {
		return count(root)
	}

	return Max(largestBSTSubtree(root.Left), largestBSTSubtree(root.Right))
}
