package golang

import "math"

func isEvenOddTree(root *TreeNode) bool {
	q := NewQueue[*TreeNode]()
	prev := &TreeNode{Val: 0, Left: nil, Right: nil}
	isEven := true

	for q.Enqueue(root); !q.IsEmpty(); {
		size := q.Size()

		if isEven {
			prev.Val = math.MinInt
		} else {
			prev.Val = math.MaxInt
		}

		for i := 0; i < size; i++ {
			curr, _ := q.Dequeue()

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
				q.Enqueue(curr.Left)
			}

			if curr.Right != nil {
				q.Enqueue(curr.Right)
			}

			prev = curr
		}

		isEven = !isEven
	}

	return true
}
