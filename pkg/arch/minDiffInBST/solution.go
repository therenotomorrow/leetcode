package minDiffInBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	val  *TreeNode
	next *Node
}

type Stack struct {
	head *Node
}

func (s *Stack) Push(val *TreeNode) {
	node := &Node{val: val}

	if s.head != nil {
		node.next = s.head
	}

	s.head = node
}

func (s *Stack) Pop() *TreeNode {
	if s.head == nil {
		return nil
	}

	val := s.head.val
	s.head = s.head.next

	return val
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}

func minDiffInBST(root *TreeNode) int {
	stack := Stack{}
	val := -1
	min := 100_000

	for root != nil || !stack.IsEmpty() {
		if root != nil {
			stack.Push(root)
			root = root.Left
		} else {
			root = stack.Pop()

			if diff := root.Val - val; diff < min {
				min = diff
			}

			val = root.Val
			root = root.Right
		}
	}

	return min
}
