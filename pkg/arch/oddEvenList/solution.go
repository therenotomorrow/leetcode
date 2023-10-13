package oddEvenList

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odds, evens := head, head
	rightHead := evens.Next

	for curr, cnt := head.Next, 1; curr != nil; curr, cnt = curr.Next, cnt+1 {
		if cnt%2 == 0 {
			odds.Next = curr
			odds = curr
		} else {
			evens.Next = curr
			evens = curr
		}
	}

	evens.Next = nil
	odds.Next = rightHead

	return head
}
