package golang

func reorderList(head *ListNode) {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	var prev *ListNode

	for curr := slow; curr != nil; {
		curr.Next, prev, curr = prev, curr, curr.Next
	}

	first := head
	second := prev

	for second.Next != nil {
		first.Next, first = second, first.Next
		second.Next, second = first, second.Next
	}
}
