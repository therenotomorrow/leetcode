package minDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	depth := 1

	switch {
	case root == nil:
		return 0
	case root.Left == nil:
		return depth + minDepth(root.Right)
	case root.Right == nil:
		return depth + minDepth(root.Left)
	}

	l := minDepth(root.Left)
	r := minDepth(root.Right)

	if l < r {
		depth += l
	} else {
		depth += r
	}

	return depth
}
