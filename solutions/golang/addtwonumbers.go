package golang

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	curr := new(ListNode)
	head := curr

	for sum := 0; l1 != nil || l2 != nil; sum = 0 {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry
		carry = sum / 10 // max = 1 + 9 + 9, so carry 0 or 1

		curr.Next = &ListNode{Val: sum % 10, Next: nil}
		curr = curr.Next
	}

	if carry != 0 {
		curr.Next = &ListNode{Val: carry, Next: nil}
	}

	return head.Next
}
