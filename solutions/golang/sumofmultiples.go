package golang

func sumOfMultiples(n int) int {
	const (
		three = 3
		five  = 5
		seven = 7
	)

	sum := 0

	for num := 1; num <= n; num++ {
		if num%three == 0 || num%five == 0 || num%seven == 0 {
			sum += num
		}
	}

	return sum
}
