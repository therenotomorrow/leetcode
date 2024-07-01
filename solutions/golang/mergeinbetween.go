package golang

func mergeInBetween(list1 *ListNode, aaa int, bbb int, list2 *ListNode) *ListNode {
	idx := 0
	head := list1
	curr := &ListNode{Val: 0, Next: list1}

	for idx != aaa {
		curr = curr.Next
		idx++
	}

	tail := curr.Next

	for idx != bbb {
		tail = tail.Next
		idx++
	}

	curr.Next = list2

	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = tail.Next

	return head
}
