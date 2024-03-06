package golang

func hasCycle(head *ListNode) bool {
	// Floyd's Cycle Finding Algorithm
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil && slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow == fast
}
