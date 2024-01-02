package swapNodes

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	slow, fast := head, head

	for i := 1; i < k; i++ {
		fast = fast.Next
	}
	ln := fast

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	rn := slow

	ln.Val, rn.Val = rn.Val, ln.Val

	return head
}
