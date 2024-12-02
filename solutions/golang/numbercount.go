package golang

func numberCount(left int, right int) int {
	isUniq := func(num int) bool {
		seen := make(map[int]struct{})

		for ; num > 0; num /= Digits {
			digit := num % Digits

			if _, ok := seen[digit]; ok {
				return false
			}

			seen[digit] = struct{}{}
		}

		return true
	}

	cnt := 0

	for num := left; num <= right; num++ {
		if isUniq(num) {
			cnt++
		}
	}

	return cnt
}
