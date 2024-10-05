package golang

func lexicalOrder(limit int) []int {
	var (
		order    = make([]int, 0, limit+1)
		populate func(currNum int)
	)

	populate = func(currNum int) {
		if currNum > limit {
			return
		}

		order = append(order, currNum)

		for digit := range Digits {
			nextNum := currNum*Digits + digit

			if nextNum > limit {
				break
			}

			populate(nextNum)
		}
	}

	for i := 1; i < Digits; i++ {
		populate(i)
	}

	return order
}
