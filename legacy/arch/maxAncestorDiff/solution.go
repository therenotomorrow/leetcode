package maxAncestorDiff

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mini(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxi(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxAncestorDiff(root *TreeNode) int {
	return maxAncestorDiffCalc(root, root.Val, root.Val)
}

func maxAncestorDiffCalc(node *TreeNode, min int, max int) int {
	if node == nil {
		return max - min
	}

	min = mini(min, node.Val)
	max = maxi(max, node.Val)

	left := maxAncestorDiffCalc(node.Left, min, max)
	right := maxAncestorDiffCalc(node.Right, min, max)

	return maxi(left, right)
}
