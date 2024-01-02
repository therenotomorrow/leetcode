package middleNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	mid := head

	for curr := head; curr != nil && curr.Next != nil; curr = curr.Next.Next {
		mid = mid.Next
	}

	return mid
}
