package getDecimalValue

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	val := head.Val

	for head.Next != nil {
		val = 2*val + head.Next.Val
		head = head.Next
	}

	return val
}
