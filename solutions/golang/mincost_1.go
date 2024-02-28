package golang

func minCost1(colors string, neededTime []int) int {
	var left, right, cost int

	for left < len(colors) && right < len(colors) {
		var totalWindowCost, maxWindowCost int

		for right < len(colors) && colors[left] == colors[right] {
			totalWindowCost += neededTime[right]
			maxWindowCost = Max(maxWindowCost, neededTime[right])
			right++
		}

		cost += totalWindowCost - maxWindowCost
		left = right
	}

	return cost
}
