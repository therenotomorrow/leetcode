package golang

func getDecimalValue(head *ListNode) int {
	val := head.Val

	for head.Next != nil {
		val = Double*val + head.Next.Val
		head = head.Next
	}

	return val
}
