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

func twoSum2(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for l < r {
		sum := numbers[l] + numbers[r]

		switch {
		case sum == target:
			return []int{l + 1, r + 1}
		case sum > target:
			r--
		default:
			l++
		}
	}

	// impossible by task description
	return []int{-1, -1}
}
