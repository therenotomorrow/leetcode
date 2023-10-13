package deleteNode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func successor(node *TreeNode) int {
	for node = node.Right; node.Left != nil; node = node.Left {
	}
	return node.Val
}

func predecessor(node *TreeNode) int {
	for node = node.Left; node.Right != nil; node = node.Right {
	}
	return node.Val
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Left == nil {
			root.Val = successor(root)
			root.Right = deleteNode(root.Right, root.Val)
		} else {
			root.Val = predecessor(root)
			root.Left = deleteNode(root.Left, root.Val)
		}
	}

	return root
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode2(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
