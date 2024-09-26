package golang

func splitListToParts(head *ListNode, k int) []*ListNode {
	size := 0
	for curr := head; curr != nil; curr = curr.Next {
		size++
	}

	remainder := size % k
	size /= k

	prev, curr := head, head
	parts := make([]*ListNode, k)

	for i := range k {
		part := curr
		partSize := size

		if remainder > 0 {
			partSize++
			remainder--
		}

		for range partSize {
			prev, curr = curr, curr.Next
		}

		if prev != nil {
			prev.Next = nil
		}

		parts[i] = part
	}

	return parts
}
