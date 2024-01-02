package longestZigZag

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type direction = int

const (
	left = iota
	right
)

func longestZigZag(root *TreeNode) int {
	var maxi int

	stepZigZag(root, left, 0, &maxi)
	stepZigZag(root, right, 0, &maxi)

	return maxi
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func stepZigZag(node *TreeNode, dir direction, steps int, maxi *int) {
	if node == nil {
		return
	}

	*maxi = max(*maxi, steps)

	switch dir {
	case left:
		stepZigZag(node.Left, right, steps+1, maxi)
		stepZigZag(node.Right, left, 1, maxi)
	case right:
		stepZigZag(node.Left, right, 1, maxi)
		stepZigZag(node.Right, left, steps+1, maxi)
	}
}
