package swapPairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dummy, prev, next *ListNode

	dummy = head.Next
	for head != nil && head.Next != nil {
		if prev != nil {
			prev.Next = head.Next
		}

		prev = head
		next = head.Next.Next
		head.Next.Next = head

		head.Next = next
		head = next
	}

	return dummy
}
