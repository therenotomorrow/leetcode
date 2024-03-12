package golang

func removeZeroSumSublists(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	prefixSum := 0
	prefixMap := make(map[int]*ListNode)

	for curr := dummy; curr != nil; curr = curr.Next {
		prefixSum += curr.Val

		prev, ok := prefixMap[prefixSum]
		if !ok {
			prefixMap[prefixSum] = curr

			continue
		}

		// find where sum is zero, we sure about it because found it in map
		curr = prev.Next

		for windowSum := prefixSum + curr.Val; windowSum != prefixSum; windowSum += curr.Val {
			delete(prefixMap, windowSum)

			curr = curr.Next
		}

		prev.Next = curr.Next
	}

	return dummy.Next
}
