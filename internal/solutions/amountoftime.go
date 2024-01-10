package solutions

import (
	"cmp"

	"github.com/therenotomorrow/leetcode/internal/structs"
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
		*time = Max(left, right)

		return -1 // signal outside that we found infected node

	case left > -1 && right > -1:
		return Max(left, right) + 1

	default:
		*time = Max(*time, Abs(left)+Abs(right))

		return Min(left, right) - 1 // how far infected node from root
	}
}

// Max returns maximum of `nums` of type `T` and have more friendly interface then build-in `max()`.
// Panics if `nums` have zero elements.
func Max[T cmp.Ordered](nums ...T) T {
	m := nums[0]

	for _, n := range nums {
		if n > m {
			m = n
		}
	}

	return m
}

// Min returns minimum of `nums` of type `T` and have more friendly interface then build-in `min()`.
// Panics if `nums` have zero elements.
func Min[T cmp.Ordered](nums ...T) T {
	m := nums[0]

	for _, n := range nums {
		if n < m {
			m = n
		}
	}

	return m
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
