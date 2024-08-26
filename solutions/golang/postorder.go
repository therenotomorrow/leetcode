package golang

func postorder(root *Node) []int {
	ans := make([]int, 0)

	if root == nil {
		return ans
	}

	for _, child := range root.Children {
		ans = append(ans, postorder(child)...)
	}

	return append(ans, root.Val)
}
