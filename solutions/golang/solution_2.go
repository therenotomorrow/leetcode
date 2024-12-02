package golang

func solution2(knows func(a int, b int) bool) func(n int) int {
	check := func(celebrity int, n int) int {
		for person := range n {
			if person == celebrity {
				continue
			}

			if knows(celebrity, person) || !knows(person, celebrity) {
				return -1
			}
		}

		return celebrity
	}

	return func(n int) int {
		celebrity := 0

		for person := 1; person < n; person++ {
			if knows(celebrity, person) {
				celebrity = person
			}
		}

		return check(celebrity, n)
	}
}
