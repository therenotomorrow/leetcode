package golang

func countPalindromicPaths(root *TreeNode, path int) int {
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

func pseudoPalindromicPaths(root *TreeNode) int {
	return countPalindromicPaths(root, 0)
}
