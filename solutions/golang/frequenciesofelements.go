package golang

func frequenciesOfElements(head *ListNode) *ListNode {
	cnt := make(map[int]int)
	for curr := head; curr != nil; curr = curr.Next {
		cnt[curr.Val]++
	}

	node := &ListNode{Val: -1, Next: nil}
	dummy := node

	for _, freq := range cnt {
		node.Next = &ListNode{Val: freq, Next: nil}
		node = node.Next
	}

	return dummy.Next
}
