package canMakeArithmeticProgression

func canMakeArithmeticProgression(arr []int) bool {
	min := arr[0]
	max := arr[0]

	for _, num := range arr {
		if num > max {
			max = num
		}

		if num < min {
			min = num
		}
	}

	if (max-min)%(len(arr)-1) != 0 {
		return false
	}

	diff := (max - min) / (len(arr) - 1)

	for i := 0; i < len(arr); {
		switch {
		case arr[i] == min+i*diff:
			i++

		case (arr[i]-min)%diff != 0:
			return false

		default:
			j := (arr[i] - min) / diff

			// duplicates
			if arr[i] == arr[j] {
				return false
			}

			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	return true
}
