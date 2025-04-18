package golang

func majorityElement(nums []int) int {
	times := len(nums) / Half
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++

		if count[num] > times {
			return num
		}
	}

	return 0
}
