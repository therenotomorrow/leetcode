package deleteMiddle

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}

	var slow, fast, last *ListNode

	slow, fast = head, head
	for fast != nil && fast.Next != nil {
		last = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	last.Next = slow.Next

	return head
}
