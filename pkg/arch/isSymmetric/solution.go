package isSymmetric

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isSame(root, root)
}

func isSame(n1 *TreeNode, n2 *TreeNode) bool {
	switch {
	case n1 == nil && n2 == nil:
		return true
	case n1 == nil || n2 == nil:
		return false
	}

	return n1.Val == n2.Val && isSame(n1.Right, n2.Left) && isSame(n1.Left, n2.Right)
}
