package reverseBetween

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// no need reverse for same indexes
	if left == right || head == nil {
		return head
	}

	curr, prev := head, head
	for i := 0; i < left-1; i++ {
		prev = curr
		curr = curr.Next
	}

	endNode := curr
	startNode := prev

	for i := left; i <= right; i++ {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	if left == 1 {
		head = prev
	} else {
		startNode.Next = prev
	}

	endNode.Next = curr

	return head
}
