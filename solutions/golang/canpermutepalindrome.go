package golang

func canPermutePalindrome(str string) bool {
	const div = 2

	cnt := make(map[rune]int, len(str))
	for _, r := range str {
		cnt[r]++
	}

	odds := 0

	for _, times := range cnt {
		odds += times % div
	}

	return odds <= 1 // 1 if odd string and 0 if even string
}
