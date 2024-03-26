package firstMissingPositive

func firstMissingPositive(nums []int) int {
	seen := make([]bool, len(nums)+1)

	for _, num := range nums {
		if num > 0 && num <= len(nums) {
			seen[num] = true
		}
	}

	for i := 1; i <= len(nums); i++ {
		if !seen[i] {
			return i
		}
	}

	return len(nums) + 1
}
