package golang

func addTwoNumbers(ln1 *ListNode, ln2 *ListNode) *ListNode {
	carry := 0
	curr := new(ListNode)
	head := curr

	for sum := 0; ln1 != nil || ln2 != nil; sum = 0 {
		if ln1 != nil {
			sum += ln1.Val
			ln1 = ln1.Next
		}

		if ln2 != nil {
			sum += ln2.Val
			ln2 = ln2.Next
		}

		sum += carry
		carry = sum / Digits // max = 1 + 9 + 9, so carry 0 or 1

		curr.Next = &ListNode{Val: sum % Digits, Next: nil}
		curr = curr.Next
	}

	if carry != 0 {
		curr.Next = &ListNode{Val: carry, Next: nil}
	}

	return head.Next
}
