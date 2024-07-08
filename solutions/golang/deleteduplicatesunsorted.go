package golang

func deleteDuplicatesUnsorted(head *ListNode) *ListNode {
	cnt := make(map[int]int)
	for curr := head; curr != nil; curr = curr.Next {
		cnt[curr.Val]++
	}

	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy

	for curr := dummy.Next; curr != nil; curr = curr.Next {
		if cnt[curr.Val] > 1 {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
	}

	return dummy.Next
}
