package removeElements

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	sent := &ListNode{Val: -1, Next: head}

	for curr, prev := head, sent; curr != nil; curr = curr.Next {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
	}

	return sent.Next
}
