package golang

import "strconv"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	const double = 2

	var (
		cnt = make(map[string]int)
		ans = make([]*TreeNode, 0)
		dfs func(root *TreeNode) string
	)

	dfs = func(root *TreeNode) string {
		if root == nil {
			return ""
		}

		key := "(" + dfs(root.Left) + strconv.Itoa(root.Val) + dfs(root.Right) + ")"

		cnt[key]++
		if cnt[key] == double {
			ans = append(ans, root)
		}

		return key
	}

	dfs(root)

	return ans
}
