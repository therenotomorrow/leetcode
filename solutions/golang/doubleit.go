package golang

func doubleIt(head *ListNode) *ListNode {
	var dynamic func(curr *ListNode) int

	dynamic = func(curr *ListNode) int {
		if curr == nil {
			return 0
		}

		val := 2*curr.Val + dynamic(curr.Next)

		curr.Val = val % 10

		return val / 10
	}

	reminder := dynamic(head)
	if reminder > 0 {
		head = &ListNode{Val: reminder, Next: head}
	}

	return head
}
