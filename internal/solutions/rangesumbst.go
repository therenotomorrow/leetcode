package solutions

import "github.com/therenotomorrow/leetcode/internal/structs"

func rangeSumBST(root *structs.TreeNode, low int, high int) int {
	sum := 0

	if root == nil {
		return sum
	}

	if low <= root.Val && root.Val <= high {
		sum += root.Val
	}

	if low < root.Val {
		sum += rangeSumBST(root.Left, low, high)
	}

	if root.Val < high {
		sum += rangeSumBST(root.Right, low, high)
	}

	return sum
}
