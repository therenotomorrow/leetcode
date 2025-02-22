package golang

func recoverFromPreorder(traversal string) *TreeNode {
	maxiIndex := len(traversal)
	currNodes := make([]*TreeNode, 0) // current nodes on depth level

	for currIndex := 0; currIndex < maxiIndex; {
		level := 0

		for currIndex < maxiIndex && traversal[currIndex] == '-' {
			level++
			currIndex++
		}

		node := &TreeNode{Val: 0, Left: nil, Right: nil}

		for currIndex < maxiIndex && traversal[currIndex] != '-' {
			node.Val = Digits*node.Val + int(traversal[currIndex]-'0')
			currIndex++
		}

		if level < len(currNodes) {
			currNodes[level] = node
		} else {
			currNodes = append(currNodes, node)
		}

		// for root node we don't need to search for parent
		if level == 0 {
			continue
		}

		parent := currNodes[level-1]
		if parent.Left == nil {
			parent.Left = node
		} else {
			parent.Right = node
		}
	}

	return currNodes[0]
}
