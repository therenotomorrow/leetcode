package golang

func isSubPath(head *ListNode, root *TreeNode) bool {
	var dfs func(head *ListNode, root *TreeNode) bool

	if root == nil {
		return false
	}

	dfs = func(head *ListNode, root *TreeNode) bool {
		if head == nil {
			return true
		}

		if root == nil || root.Val != head.Val {
			return false
		}

		return dfs(head.Next, root.Left) || dfs(head.Next, root.Right)
	}

	return dfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}
