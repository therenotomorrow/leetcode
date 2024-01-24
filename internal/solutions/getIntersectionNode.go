package solutions

import "github.com/therenotomorrow/leetcode/internal/structs"

func getIntersectionNode(headA *structs.ListNode, headB *structs.ListNode) *structs.ListNode {
	nodes := make(map[*structs.ListNode]struct{})

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
