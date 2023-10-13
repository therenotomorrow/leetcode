package fastAndSlowPointers

type ListNode struct {
	val  int
	next *ListNode
}

func getMiddle(head *ListNode) int {
	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow.val
}

func findNode(head *ListNode, k int) *ListNode {
	slow := head
	fast := head

	for i := 0; i < k; i++ {
		fast = fast.next
	}

	for fast != nil {
		slow = slow.next
		fast = fast.next
	}

	return slow
}
