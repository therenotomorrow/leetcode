package golang

func createBinaryTree(descriptions [][]int) *TreeNode {
	tree := make(map[int]*TreeNode)
	vals := make(map[int]struct{})

	for _, description := range descriptions {
		parent := &TreeNode{Val: description[0], Left: nil, Right: nil}
		child := &TreeNode{Val: description[1], Left: nil, Right: nil}

		if _, ok := tree[child.Val]; !ok {
			tree[child.Val] = child
		} else {
			child = tree[child.Val]
		}

		if _, ok := tree[parent.Val]; !ok {
			tree[parent.Val] = parent
		} else {
			parent = tree[parent.Val]
		}

		if description[2] == 1 {
			parent.Left = child
		} else {
			parent.Right = child
		}

		vals[child.Val] = struct{}{}
	}

	root := &TreeNode{Val: 0, Left: nil, Right: nil}

	for _, parent := range tree {
		if _, ok := vals[parent.Val]; !ok {
			root = parent

			break
		}
	}

	return root
}
