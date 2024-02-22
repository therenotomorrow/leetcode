package golang

func minCostClimbingStairs(cost []int) int {
	steps := make([]int, len(cost)+1)

	for i := 2; i < len(cost)+1; i++ {
		steps[i] = Min(steps[i-1]+cost[i-1], steps[i-2]+cost[i-2])
	}

	return steps[len(cost)]
}
