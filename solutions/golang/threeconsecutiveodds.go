package golang

func threeConsecutiveOdds(arr []int) bool {
	const size = 3

	if len(arr) < size {
		return false
	}

	cnt := 0

	for _, num := range arr {
		if num%Half != 0 {
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
