package golang

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	dummy := head
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}

		return a
	}

	for curr := head; curr != nil && curr.Next != nil; {
		node := &ListNode{Val: gcd(curr.Val, curr.Next.Val), Next: curr.Next}
		curr.Next = node
		curr = node.Next
	}

	return dummy
}
