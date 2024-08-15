package golang

func lemonadeChange(bills []int) bool {
	const (
		five   = 5
		ten    = 10
		twenty = 20
	)

	var (
		fives = 0
		tens  = 0
	)

	for _, bill := range bills {
		switch bill {
		case five:
			fives++
		case ten:
			fives--
			tens++
		case twenty:
			if tens > 0 {
				fives--
				tens--
			} else {
				fives -= 3
			}
		}

		if fives < 0 || tens < 0 {
			return false
		}
	}

	return true
}
