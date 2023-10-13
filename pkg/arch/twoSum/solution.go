package twoSum

func twoSum(nums []int, target int) []int {
	diff := make(map[int]int)

	for i, num := range nums {
		j, ok := diff[target-num]
		if ok {
			return []int{i, j}
		}

		diff[num] = i
	}

	return nums
}
