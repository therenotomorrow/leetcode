package deleteDuplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	ansHead := head
	curr := ansHead

	for head != nil {
		if curr.Val != head.Val {
			curr.Next = head
			curr = curr.Next
		}

		head = head.Next
	}

	curr.Next = nil

	return ansHead
}

func deleteDuplicates2(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -101, Next: head}

	slow, fast := dummy, head
	for fast != nil {
		if fast.Next != nil && fast.Val == fast.Next.Val {
			for fast.Next != nil && fast.Val == fast.Next.Val {
				fast = fast.Next
			}

			slow.Next = fast.Next
		} else {
			slow = slow.Next
		}

		fast = fast.Next
	}

	return dummy.Next
}
