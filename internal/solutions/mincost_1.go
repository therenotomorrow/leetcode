package solutions

import "github.com/therenotomorrow/leetcode/pkg/mathfunc"

func minCost1(colors string, neededTime []int) int {
	left := 0
	right := 0
	cost := 0

	for left < len(colors) && right < len(colors) {
		totalWindowCost := 0
		maxWindowCost := 0

		for right < len(colors) && colors[left] == colors[right] {
			totalWindowCost += neededTime[right]
			maxWindowCost = mathfunc.Max(maxWindowCost, neededTime[right])
			right++
		}

		cost += totalWindowCost - maxWindowCost
		left = right
	}

	return cost
}
