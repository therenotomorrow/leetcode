package golang

func treeToDoublyList(root *TreeNode) *TreeNode {
	var (
		dfs  func(node *TreeNode)
		head *TreeNode
		tail *TreeNode
	)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)

		if tail != nil {
			node.Left, tail.Right = tail, node
		} else {
			head = node
		}

		tail = node

		dfs(node.Right)
	}

	dfs(root)

	if tail != nil {
		tail.Right = head
	}

	if head != nil {
		head.Left = tail
	}

	return head
}
