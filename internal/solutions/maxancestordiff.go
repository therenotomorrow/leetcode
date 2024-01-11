package solutions

import (
	"github.com/therenotomorrow/leetcode/internal/structs"
	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
)

func maxAncestorDiff(root *structs.TreeNode) int {
	return maxAncestorDiffCalc(root, root.Val, root.Val)
}

func maxAncestorDiffCalc(node *structs.TreeNode, min int, max int) int {
	if node == nil {
		return max - min
	}

	min = mathfunc.Min(min, node.Val)
	max = mathfunc.Max(max, node.Val)

	lDiff := maxAncestorDiffCalc(node.Left, min, max)
	rDiff := maxAncestorDiffCalc(node.Right, min, max)

	return mathfunc.Max(lDiff, rDiff)
}
