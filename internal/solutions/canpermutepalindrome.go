package solutions

func canPermutePalindrome(s string) bool {
	cnt := make(map[rune]int, len(s))
	for _, r := range s {
		cnt[r]++
	}

	odds := 0

	for _, times := range cnt {
		odds += times % 2
	}

	return odds <= 1 // 1 if odd string and 0 if even string
}
