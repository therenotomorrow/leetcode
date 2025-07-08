package golang

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var curr *TreeNode

	for root != nil {
		if p.Val < root.Val {
			curr = root
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return curr
}
