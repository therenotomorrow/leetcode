package golang

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow := head
	fast := head

	for range n {
		fast = fast.Next
	}

	var last *ListNode
	for fast != nil {
		last = slow
		slow = slow.Next
		fast = fast.Next
	}

	if last == nil {
		head = slow.Next
	} else {
		last.Next = slow.Next
	}

	return head
}
