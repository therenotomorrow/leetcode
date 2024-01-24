package solutions

func isStrobogrammatic(num string) bool {
	left := 0
	right := len(num) - 1
	mirrors := map[byte]byte{'0': '0', '1': '1', '6': '9', '8': '8', '9': '6'}

	for left <= right {
		mustBe, ok := mirrors[num[left]]
		if !ok || mustBe != num[right] {
			return false
		}

		right--
		left++
	}

	return true
}
