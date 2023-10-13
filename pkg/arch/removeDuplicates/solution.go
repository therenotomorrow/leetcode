package removeDuplicates

func removeDuplicates(s string) string {
	var stack = make([]rune, 0)

	for _, r := range s {
		if l := len(stack); l > 0 && stack[l-1] == r {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}

	return string(stack)
}

func removeDuplicates2(nums []int) int {
	var i, j int

	for i = 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}
