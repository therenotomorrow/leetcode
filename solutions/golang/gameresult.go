package golang

func gameResult(head *ListNode) string {
	odds := 0
	evens := 0

	for head != nil {
		if head.Next.Val > head.Val {
			odds++
		} else {
			evens++
		}

		head = head.Next.Next
	}

	switch {
	case odds > evens:
		return "Odd"
	case evens > odds:
		return "Even"
	default:
		return "Tie"
	}
}
