package golang

func reverseList(head *ListNode) *ListNode {
	curr := head
	stack := NewStack[*ListNode]()

	for ; curr != nil; curr = curr.Next {
		stack.Push(curr)
	}

	head = &ListNode{Val: 0, Next: nil}
	curr = head

	for ; !stack.IsEmpty(); curr = curr.Next {
		node, _ := stack.Pop()
		curr.Next = node
	}

	curr.Next = nil

	return head.Next
}
