package golang

func collectTree(root *TreeNode) []int {
	stack := NewStack[*TreeNode]()
	arr := make([]int, 0)

	for ok := true; ok; {
		if root != nil {
			stack.Push(root)
			root = root.Left
		} else {
			root, ok = stack.Pop()
			if !ok {
				break
			}

			if root.Left == nil && root.Right == nil {
				arr = append(arr, root.Val)
			}

			root = root.Right
		}
	}

	return arr
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	arr1 := collectTree(root1)
	arr2 := collectTree(root2)

	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
