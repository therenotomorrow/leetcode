package golang

func missingNumber(nums []int) int {
	set := NewSet[int]()

	for _, num := range nums {
		set.Add(num)
	}

	for num := 0; num <= len(nums)+1; num++ {
		if !set.Contains(num) {
			return num
		}
	}

	// impossible due task description
	return 0
}
