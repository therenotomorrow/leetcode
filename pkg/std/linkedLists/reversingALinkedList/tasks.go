package reversingALinkedList

type ListNode struct {
	val  int
	next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev, curr *ListNode

	curr = head
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	return prev
}
