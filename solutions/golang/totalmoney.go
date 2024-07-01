package golang

func totalMoney(n int) int {
	const (
		days   = 7
		double = 2
		half   = double
	)

	a := n / days // cycles
	b := n % days // rest days

	// arithmetic sequence reorganized
	return (days*a*(a+days) + double*a*b + b*(b+1)) / half
}
