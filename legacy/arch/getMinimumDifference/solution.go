package getMinimumDifference

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	arr := collectNodes(root)

	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return root.Val
	}

	min := 100_000
	for i := 0; i < len(arr)-1; i++ {
		if diff := arr[i+1] - arr[i]; diff < min {
			min = diff
		}
	}

	return min
}

func collectNodes(node *TreeNode) []int {
	arr := make([]int, 0)

	if node.Left != nil {
		arr = append(arr, collectNodes(node.Left)...)
	}

	arr = append(arr, node.Val)

	if node.Right != nil {
		arr = append(arr, collectNodes(node.Right)...)
	}

	return arr
}
