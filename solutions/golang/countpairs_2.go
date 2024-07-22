package golang

func countPairs2(root *TreeNode, distance int) int {
	var (
		cnt int
		dfs func(node *TreeNode) []int
	)

	dfs = func(node *TreeNode) []int {
		if node == nil {
			return []int{}
		}

		if node.Left == nil && node.Right == nil {
			return []int{1}
		}

		leftDistances := dfs(node.Left)
		rightDistances := dfs(node.Right)

		for _, left := range leftDistances {
			for _, right := range rightDistances {
				if left+right <= distance {
					cnt++
				}
			}
		}

		distances := make([]int, 0)

		for _, distRange := range [][]int{leftDistances, rightDistances} {
			for _, dist := range distRange {
				if dist+1 <= distance {
					distances = append(distances, dist+1)
				}
			}
		}

		return distances
	}

	_ = dfs(root)

	return cnt
}
