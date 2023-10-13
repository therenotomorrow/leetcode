package diameterOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var diameter int

func maxi(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func diameterOfBinaryTree(root *TreeNode) int {
	diameter = 0
	findPath(root)
	return diameter
}

func findPath(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftPath := findPath(node.Left)
	rightPath := findPath(node.Right)

	diameter = maxi(diameter, leftPath+rightPath)

	return maxi(leftPath, rightPath) + 1
}
