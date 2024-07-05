package golang

func nodesBetweenCriticalPoints(head *ListNode) []int {
	const minPoints = 2

	ans := []int{-1, -1}
	prev := head
	curr := head.Next
	dist := make([]int, 0)

	for i := 2; curr.Next != nil; i++ {
		if prev.Val < curr.Val && curr.Val > curr.Next.Val {
			dist = append(dist, i)
		} else if prev.Val > curr.Val && curr.Val < curr.Next.Val {
			dist = append(dist, i)
		}

		prev = curr
		curr = curr.Next
	}

	// we need more than 1 critical point
	if len(dist) < minPoints {
		return ans
	}

	ans[1] = dist[len(dist)-1] - dist[0]
	ans[0] = ans[1]

	for i := range len(dist) - 1 {
		ans[0] = Min(ans[0], dist[i+1]-dist[i])
	}

	return ans
}
