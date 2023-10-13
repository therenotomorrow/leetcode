package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		carry int
		curr  = new(ListNode)
		head  = curr
	)

	for l1 != nil || l2 != nil {
		var val1 int
		if l1 != nil {
			val1, l1 = l1.Val, l1.Next
		}

		var val2 int
		if l2 != nil {
			val2, l2 = l2.Val, l2.Next
		}

		sum := carry + val1 + val2
		carry = sum / 10 // max = 1 + 9 + 9, so carry 0 or 1

		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}

	if carry != 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return head.Next
}
