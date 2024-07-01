package golang

func majorityElement(nums []int) int {
	const half = 2

	times := len(nums) / half
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++

		if count[num] > times {
			return num
		}
	}

	return 0
}
