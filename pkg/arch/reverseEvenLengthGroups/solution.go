package reverseEvenLengthGroups

type ListNode struct {
	Val  int
	Next *ListNode
}

func swap(arr []int) []int {
	for i, j := 0, len(arr)-1; i <= j; {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	oldHead := head
	forward := make([]int, 0)
	reversed := make([]int, 0)

	cnt := 0
	limit := 1
	reverse := false
	groups := make([][]int, 0)

	for head != nil {
		if cnt == limit {
			if cnt%2 == 0 {
				groups = append(groups, swap(reversed))
				reversed = make([]int, 0)
			} else {
				groups = append(groups, forward)
				forward = make([]int, 0)
			}

			reverse = !reverse
			cnt = 1
			limit++
		} else {
			cnt++
		}

		if reverse {
			reversed = append(reversed, head.Val)
		} else {
			forward = append(forward, head.Val)
		}

		head = head.Next
	}

	if l := len(reversed); l > 0 && l%2 == 0 {
		groups = append(groups, swap(reversed))
	}
	if l := len(forward); l > 0 && l%2 == 0 {
		groups = append(groups, swap(forward))
	}

	oldHead2 := oldHead

	for _, group := range groups {
		for _, item := range group {
			oldHead.Val = item
			oldHead = oldHead.Next
		}
	}

	return oldHead2
}
