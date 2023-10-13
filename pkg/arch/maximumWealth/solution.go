package maximumWealth

func sum(nums []int) (s int) {
	for _, n := range nums {
		s += n
	}
	return
}

func maximumWealth(accounts [][]int) (max int) {
	for _, acc := range accounts {
		tmp := sum(acc)
		if tmp > max {
			max = tmp
		}
	}

	return
}
