package golang

func averageWaitingTime(customers [][]int) float64 {
	finished := 0
	wait := float64(0)

	for _, customer := range customers {
		finished = Max(finished, customer[0]) + customer[1]

		wait += float64(finished - customer[0])
	}

	return wait / float64(len(customers))
}
