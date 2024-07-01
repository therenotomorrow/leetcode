package golang

func findBottomLeftValue(root *TreeNode) int {
	que := NewQueue[*TreeNode]()
	ans := 0

	for que.Enqueue(root); !que.IsEmpty(); {
		size := que.Size()
		peek, _ := que.Peek()
		ans = peek.Val

		for range size {
			node, _ := que.Dequeue()

			if node.Left != nil {
				que.Enqueue(node.Left)
			}

			if node.Right != nil {
				que.Enqueue(node.Right)
			}
		}
	}

	return ans
}
