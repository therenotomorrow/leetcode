package solutions

import "github.com/therenotomorrow/leetcode/internal/structs"

func countPalindromicPaths(root *structs.TreeNode, path int) int {
	count := 0

	if root == nil {
		return count
	}

	// tricky bitwise
	path ^= 1 << root.Val

	if root.Left == nil && root.Right == nil {
		if path&(path-1) == 0 {
			count++
		}
	} else {
		count += countPalindromicPaths(root.Left, path)
		count += countPalindromicPaths(root.Right, path)
	}

	return count
}

func pseudoPalindromicPaths(root *structs.TreeNode) int {
	return countPalindromicPaths(root, 0)
}
