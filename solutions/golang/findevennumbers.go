package golang

func findEvenNumbers(digits []int) []int {
	const (
		minNum = 100
		maxNum = 1000
	)

	cnt := make(map[int]int)

	for _, digit := range digits {
		cnt[digit]++
	}

	ans := make([]int, 0)

	for num := minNum; num < maxNum; num += 2 {
		curr := make(map[int]int)

		curr[num/Digits/Digits]++
		curr[num/Digits%Digits]++
		curr[num%Digits]++

		save := true

		for digit, times := range curr {
			if cnt[digit] < times {
				save = false
			}
		}

		if save {
			ans = append(ans, num)
		}
	}

	return ans
}
