package golang

func findErrorNums(nums []int) []int {
	cnt := make(map[int]int, len(nums))

	for _, num := range nums {
		cnt[num]++
	}

	var missing, duplicate int

	for i := 1; i <= len(nums); i++ {
		times, ok := cnt[i]

		if !ok {
			missing = i
		}

		if times > 1 {
			duplicate = i
		}
	}

	return []int{duplicate, missing}
}
