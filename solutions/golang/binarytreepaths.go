package golang

import (
	"strconv"
)

func preOrderTraversal(root *TreeNode, curr string, paths []string) []string {
	if root == nil {
		return paths
	}

	curr += strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		paths = append(paths, curr)
	} else {
		curr += "->"
		paths = preOrderTraversal(root.Right, curr, preOrderTraversal(root.Left, curr, paths))
	}

	return paths
}

func binaryTreePaths(root *TreeNode) []string {
	return preOrderTraversal(root, "", make([]string, 0))
}
