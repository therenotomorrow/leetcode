package reverseList

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	stack := make([]*ListNode, 0)

	for curr := head; curr != nil; {
		stack = append(stack, curr)
		curr = curr.Next
	}

	curr := &ListNode{}
	head = curr
	for i := len(stack) - 1; i >= 0; i-- {
		curr.Next = stack[i]
		curr = curr.Next
	}
	curr.Next = nil

	return head.Next
}
