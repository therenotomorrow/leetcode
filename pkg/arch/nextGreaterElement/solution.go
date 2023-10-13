package nextGreaterElement

type Node struct {
	num  int
	next *Node
}

type Stack struct {
	head *Node
	len  int
}

func (s *Stack) Len() int {
	return s.len
}

func (s *Stack) Next() bool {
	return s.Len() > 0
}

func (s *Stack) Push(num int) {
	s.head = &Node{num: num, next: s.head}
	s.len++
}

func (s *Stack) Pop() int {
	if !s.Next() {
		return -1
	}

	node := s.head
	s.head = s.head.next
	s.len--

	return node.num
}

func (s *Stack) Peek() int {
	if !s.Next() {
		return -1
	}

	return s.head.num
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := Stack{}
	mapper := make(map[int]int)

	for _, num := range nums2 {
		for stack.Next() && num > stack.Peek() {
			mapper[stack.Pop()] = num
		}
		stack.Push(num)
	}

	for stack.Next() {
		mapper[stack.Pop()] = -1
	}

	for i, num := range nums1 {
		nums1[i] = mapper[num]
	}

	return nums1
}
