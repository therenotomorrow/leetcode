package leafSimilar

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

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val *TreeNode) {
	s.head = &Node{val: val, next: s.head}
}

func (s *Stack) Pop() *TreeNode {
	if s.head == nil {
		return nil
	}

	n := s.head
	s.head = s.head.next

	return n.val
}

func collectTree(root *TreeNode) []int {
	s := NewStack()
	arr := make([]int, 0)

	for curr := root; true; {
		if curr != nil {
			s.Push(curr)
			curr = curr.Left
		} else {
			curr = s.Pop()

			if curr == nil {
				break
			}

			if curr.Left == nil && curr.Right == nil {
				arr = append(arr, curr.Val)
			}

			curr = curr.Right
		}
	}

	return arr
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	arr1 := collectTree(root1)
	arr2 := collectTree(root2)

	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
