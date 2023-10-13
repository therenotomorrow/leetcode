package searchBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	curr := root

	for curr != nil {
		if val == curr.Val {
			return curr
		}

		if val < curr.Val {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return nil
}
