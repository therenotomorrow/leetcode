package golang

func sumSubarrayMins(arr []int) int {
	stack := NewStack[int]()
	dynamic := make([]int, len(arr))

	for i := range arr {
		for peek, ok := stack.Peek(); ok && arr[peek] >= arr[i]; peek, ok = stack.Peek() {
			stack.Pop()
		}

		if j, ok := stack.Peek(); ok {
			dynamic[i] = dynamic[j] + (i-j)*arr[i]
		} else {
			dynamic[i] = (i + 1) * arr[i]
		}

		stack.Push(i)
	}

	return Sum(dynamic...) % MOD
}
