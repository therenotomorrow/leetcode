package pairSum

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	slow := head
	fast := head

	// step 1: split list
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// step 2: reverse second
	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	// step 3: find sum
	max := 0
	sum := 0
	for prev != nil {
		sum = head.Val + prev.Val

		if sum > max {
			max = sum
		}

		head = head.Next
		prev = prev.Next
	}

	return max
}
