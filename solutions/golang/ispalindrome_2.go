package golang

func isPalindrome2(head *ListNode) bool {
	var (
		fast = head
		slow = head
		prev *ListNode
	)

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	fast = slow
	for fast != nil {
		next := fast.Next
		fast.Next = prev
		prev = fast
		fast = next
	}

	for prev != nil {
		if head.Val != prev.Val {
			return false
		}

		head = head.Next
		prev = prev.Next
	}

	return true
}
