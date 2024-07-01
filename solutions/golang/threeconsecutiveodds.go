package golang

func threeConsecutiveOdds(arr []int) bool {
	const (
		size = 3
		half = 2
	)

	if len(arr) < size {
		return false
	}

	cnt := 0

	for _, num := range arr {
		if num%half != 0 {
			cnt++
		} else {
			cnt = 0
		}

		if cnt == size {
			return true
		}
	}

	return false
}
