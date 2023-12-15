package solutions

import (
	"github.com/therenotomorrow/leetcode/internal/structs"
	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
)

func treeHeight(root *structs.TreeNode) int {
	if root == nil {
		return 0
	}

	left := treeHeight(root.Left)
	right := treeHeight(root.Right)

	return 1 + mathfunc.Max(left, right)
}

func isBalanced(root *structs.TreeNode) bool {
	if root == nil {
		return true
	}

	lh := treeHeight(root.Left)
	rh := treeHeight(root.Right)

	if mathfunc.Abs(rh-lh) < 2 && isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}

	return false
}
