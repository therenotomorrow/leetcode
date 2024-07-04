package golang

func mergeNodes(head *ListNode) *ListNode {
	curr := head.Next
	next := curr

	for curr != nil {
		sum := 0

		for next.Val != 0 {
			sum += next.Val
			next = next.Next
		}

		curr.Val = sum
		next = next.Next
		curr.Next = next
		curr = curr.Next
	}

	return head.Next
}
