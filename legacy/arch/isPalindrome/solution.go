package isPalindrome

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome2(head *ListNode) bool {
	var (
		fast = head
		slow = head
		prev *ListNode
	)

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	fast = slow
	for fast != nil {
		next := fast.Next
		fast.Next = prev
		prev = fast
		fast = next
	}

	for prev != nil {
		if head.Val != prev.Val {
			return false
		}

		head = head.Next
		prev = prev.Next
	}

	return true
}

func isPalindrome3(s string) bool {
	arr := make([]rune, 0)

	for _, r := range s {
		if 'A' <= r && r <= 'Z' {
			r += 32
		}

		if '0' <= r && r <= '9' || 'a' <= r && r <= 'z' {
			arr = append(arr, r)
		}
	}

	n := len(arr)
	for i := 0; i < n/2; i++ {
		if arr[i] != arr[n-i-1] {
			return false
		}
	}

	return true
}
