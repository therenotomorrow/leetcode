package golang

func cloneGraph(node *Node) *Node {
	var (
		visited = make(map[*Node]*Node)
		dfs     func(curr *Node) *Node
	)

	dfs = func(curNode *Node) *Node {
		if curNode == nil {
			return nil
		}

		if visited[curNode] != nil {
			return visited[curNode]
		}

		newNode := &Node{Val: curNode.Val, Children: nil}

		visited[curNode] = newNode

		for _, neighbor := range curNode.Children {
			newNode.Children = append(newNode.Children, dfs(neighbor))
		}

		return newNode
	}

	return dfs(node)
}
