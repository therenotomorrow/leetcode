package golang

func nthUglyNumber(n int) int {
	const (
		two   = 2
		three = 3
		five  = 5
	)

	if n < 0 {
		return 0
	}

	seen := NewSet[int]()
	curr := 1

	seen.Add(curr)

	for range n {
		curr = Min(seen.Values()...)

		seen.Del(curr)

		seen.Add(two * curr)
		seen.Add(three * curr)
		seen.Add(five * curr)
	}

	return curr
}
