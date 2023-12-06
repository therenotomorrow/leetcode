package solutions

func totalMoney(n int) int {
	a := n / 7 // cycles
	b := n % 7 // rest days

	// arithmetic sequence reorganized
	return (7*a*(a+7) + 2*a*b + b*(b+1)) / 2
}
