package golang

func modifiedList(nums []int, head *ListNode) *ListNode {
	set := NewSet[int]()

	set.Populate(nums...)

	dummy := &ListNode{Val: 0, Next: head}
	root := dummy

	for curr := head; curr != nil; curr = curr.Next {
		if !set.Contains(curr.Val) {
			dummy.Next = curr
			dummy = dummy.Next
		}
	}

	// remove any trailing nodes
	dummy.Next = nil

	return root.Next
}
