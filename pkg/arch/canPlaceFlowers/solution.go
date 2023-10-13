package canPlaceFlowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	for i, l := 0, len(flowerbed); i < l; i++ {
		if flowerbed[i] != 0 {
			continue
		}

		left := i == 0 || flowerbed[i-1] == 0
		right := i == l-1 || flowerbed[i+1] == 0

		if left && right {
			flowerbed[i] = 1
			n--
		}

		if n == 0 {
			return true
		}
	}

	return false
}
