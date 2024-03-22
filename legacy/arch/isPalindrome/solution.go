package isPalindrome

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
