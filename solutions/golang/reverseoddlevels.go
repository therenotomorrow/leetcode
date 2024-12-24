package golang

func reverseOddLevels(root *TreeNode) *TreeNode {
	const half = 2

	if root == nil {
		return root
	}

	var (
		que                = NewQueue[*TreeNode]()
		level, left, right int
	)

	que.Enqueue(root)

	for ; !que.IsEmpty(); level++ {
		size := que.Size()
		nodes := make([]*TreeNode, 0)

		for range size {
			curr, _ := que.Dequeue()

			if curr.Left != nil {
				que.Enqueue(curr.Left)
			}

			if curr.Right != nil {
				que.Enqueue(curr.Right)
			}

			nodes = append(nodes, curr)
		}

		if level%half == 0 {
			continue
		}

		left = 0
		right = len(nodes) - 1

		for left < right {
			nodes[left].Val, nodes[right].Val = nodes[right].Val, nodes[left].Val
			left++
			right--
		}
	}

	return root
}
