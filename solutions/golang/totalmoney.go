package golang

func totalMoney(n int) int {
	const days = 7

	a := n / days // cycles
	b := n % days // rest days

	// arithmetic sequence reorganized
	return (days*a*(a+days) + Double*a*b + b*(b+1)) / Half
}
