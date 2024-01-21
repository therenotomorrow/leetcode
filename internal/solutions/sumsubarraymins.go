package solutions

import (
	"github.com/therenotomorrow/leetcode/internal/structs"
	"github.com/therenotomorrow/leetcode/pkg/datastruct"
	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
)

func sumSubarrayMins(arr []int) int {
	stack := datastruct.NewStack[int]()
	dynamic := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
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

	return mathfunc.Sum(dynamic...) % structs.MOD
}
