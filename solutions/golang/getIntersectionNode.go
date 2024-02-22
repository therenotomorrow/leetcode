package golang

func getIntersectionNode(headA *ListNode, headB *ListNode) *ListNode {
	nodes := make(map[*ListNode]struct{})

	for curr := headA; curr != nil; curr = curr.Next {
		nodes[curr] = struct{}{}
	}

	for curr := headB; curr != nil; curr = curr.Next {
		if _, ok := nodes[curr]; ok {
			return curr
		}
	}

	return nil
}
