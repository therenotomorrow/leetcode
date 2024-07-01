package golang

import "math"

func getValue(isEven bool) int {
	if isEven {
		return math.MinInt
	}

	return math.MaxInt
}

func isEvenOddTree(root *TreeNode) bool {
	que := NewQueue[*TreeNode]()
	prev := &TreeNode{Val: 0, Left: nil, Right: nil}
	isEven := true

	for que.Enqueue(root); !que.IsEmpty(); {
		prev.Val = getValue(isEven)
		size := que.Size()

		for range size {
			curr, _ := que.Dequeue()

			if isEven == (curr.Val%2 == 0) {
				return false
			}

			if isEven && curr.Val <= prev.Val {
				return false
			}

			if !isEven && curr.Val >= prev.Val {
				return false
			}

			if curr.Left != nil {
				que.Enqueue(curr.Left)
			}

			if curr.Right != nil {
				que.Enqueue(curr.Right)
			}

			prev = curr
		}

		isEven = !isEven
	}

	return true
}
