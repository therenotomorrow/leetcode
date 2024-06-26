package golang

func countDigits(num int) int {
	cnt := 0

	for num > 0 {
		num /= Digits
		cnt++
	}

	return cnt
}

func findNumbers(nums []int) int {
	cnt := 0

	for _, num := range nums {
		if countDigits(num)%2 == 0 {
			cnt++
		}
	}

	return cnt
}
