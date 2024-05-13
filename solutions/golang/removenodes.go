package golang

func removeNodes(head *ListNode) *ListNode {
	curr := head
	currMax := 0
	stack := NewStack[*ListNode]()

	for curr != nil {
		stack.Push(curr)
		currMax = curr.Val
		curr = curr.Next
	}

	curr, _ = stack.Pop()
	head = &ListNode{Val: curr.Val, Next: nil}

	for !stack.IsEmpty() {
		curr, _ = stack.Pop()

		if curr.Val < currMax {
			continue
		}

		currMax = curr.Val
		head = &ListNode{Val: curr.Val, Next: head}
	}

	return head
}
