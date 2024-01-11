package solutions

import (
	"github.com/therenotomorrow/leetcode/internal/structs"
	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
)

func amountOfTime(root *structs.TreeNode, start int) int {
	var time int

	traversal(root, start, &time)

	return time
}

func traversal(root *structs.TreeNode, start int, time *int) int {
	if root == nil {
		return 0
	}

	left := traversal(root.Left, start, time)
	right := traversal(root.Right, start, time)

	switch {
	case root.Val == start:
		*time = mathfunc.Max(left, right)

		return -1 // signal outside that we found infected node

	case left > -1 && right > -1:
		return mathfunc.Max(left, right) + 1

	default:
		*time = mathfunc.Max(*time, mathfunc.Abs(left)+mathfunc.Abs(right))

		return mathfunc.Min(left, right) - 1 // how far infected node from root
	}
}
