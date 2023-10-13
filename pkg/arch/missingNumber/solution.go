package missingNumber

func missingNumber(nums []int) int {
	set := make([]bool, len(nums)+1)

	for _, num := range nums {
		set[num] = true
	}

	for val, ok := range set {
		if !ok {
			return val
		}
	}

	// impossible due task description
	return 0
}
