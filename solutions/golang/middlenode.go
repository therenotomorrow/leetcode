package golang

func middleNode(head *ListNode) *ListNode {
	mid := head

	for curr := head; curr != nil && curr.Next != nil; curr = curr.Next.Next {
		mid = mid.Next
	}

	return mid
}
