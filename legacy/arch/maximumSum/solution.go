package maximumSum

func maximumSum(nums []int) int {
	m := make(map[int]int)
	max := -1

	for _, num := range nums {
		key := sumDigits(num)
		val := m[key]

		if val != 0 {
			max = maximum(max, val+num)
		}

		m[key] = maximum(m[key], num)
	}

	return max
}

func sumDigits(a int) int {
	sum := 0

	for a > 0 {
		sum += a % 10
		a /= 10
	}

	return sum
}

func maximum(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
