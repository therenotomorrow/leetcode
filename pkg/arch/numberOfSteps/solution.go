package numberOfSteps

func numberOfSteps(num int) (steps int) {
	for ; num != 0; steps++ {
		if (num & 1) == 0 {
			num >>= 1
		} else {
			num--
		}
	}

	return
}
