package golang

func toArray(head *DoubleListNode) []int {
	ans := make([]int, 0)

	for ; head != nil; head = head.Next {
		ans = append(ans, head.Val)
	}

	return ans
}
