package golang

func findBottomLeftValue(root *TreeNode) int {
	q := NewQueue[*TreeNode]()
	ans := 0

	for q.Enqueue(root); !q.IsEmpty(); {
		size := q.Size()
		peek, _ := q.Peek()
		ans = peek.Val

		for i := 0; i < size; i++ {
			node, _ := q.Dequeue()

			if node.Left != nil {
				q.Enqueue(node.Left)
			}

			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}
	}

	return ans
}
