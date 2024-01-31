package solutions

import "github.com/therenotomorrow/leetcode/pkg/datastruct"

func dailyTemperatures(temperatures []int) []int {
	days := make([]int, len(temperatures))
	stack := datastruct.NewStack[int]()

	for i, temp := range temperatures {
		for j, ok := stack.Peek(); ok && temperatures[j] < temp; j, ok = stack.Peek() {
			days[j] = i - j
			stack.Pop()
		}

		stack.Push(i)
	}

	return days
}
