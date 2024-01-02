package validateStackSequences

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)

	var i int
	for _, val := range pushed {
		stack = append(stack, val)

		for len(stack) > 0 && popped[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			i++
		}
	}

	return i == len(pushed)
}
