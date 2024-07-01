package golang

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	children := NewSet[int]()

	children.Populate(leftChild...)
	children.Populate(rightChild...)

	root := -1

	for node := range n {
		// root cannot be a child node
		if !children.Contains(node) {
			root = node

			break
		}
	}

	if root == -1 {
		return false
	}

	stack := NewStack[int]()
	stack.Push(root)

	cnt := 0 // track visited nodes

	for ; !stack.IsEmpty(); cnt++ {
		node, _ := stack.Pop()

		for _, child := range [2]int{leftChild[node], rightChild[node]} {
			if child == -1 {
				continue
			}

			// already visit child before
			if !children.Contains(child) {
				return false
			}

			children.Del(child)
			stack.Push(child)
		}
	}

	return cnt == n
}
