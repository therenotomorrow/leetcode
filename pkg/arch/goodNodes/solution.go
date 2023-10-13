package goodNodes

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return goodNodesSoFar(root, math.MinInt)
}

func maxi(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func goodNodesSoFar(node *TreeNode, maxSoFar int) int {
	if node == nil {
		return 0
	}

	ans := 0
	if node.Val >= maxSoFar {
		ans++
	}

	ans += goodNodesSoFar(node.Left, maxi(node.Val, maxSoFar))
	ans += goodNodesSoFar(node.Right, maxi(node.Val, maxSoFar))

	return ans
}
