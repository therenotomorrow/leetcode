package inorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)

	arr := make([]int, 0)
	for curr := root; ; {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
			continue
		}

		if len(stack) == 0 {
			break
		}

		curr = stack[len(stack)-1]
		arr = append(arr, curr.Val)
		stack = stack[:len(stack)-1]

		curr = curr.Right
	}

	return arr
}
