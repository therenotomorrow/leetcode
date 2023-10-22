package maximumScore

type Stack struct {
	data []int
	size int
}

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
	s.size++
}

func (s *Stack) Pop() int {
	s.size--

	val := s.data[s.size]
	s.data = s.data[:s.size]

	return val
}

func (s *Stack) Peek() int {
	return s.data[s.size-1]
}

func (s *Stack) Ok() bool {
	return len(s.data) != 0
}

func (s *Stack) Clear() {
	s.data = make([]int, 0)
	s.size = 0
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumScore(nums []int, k int) int {
	// fixme: need implement MonoStack

	n := len(nums)
	stack := Stack{}

	stack.Clear()

	left := make([]int, n)
	for i := range left {
		left[i] = -1
	}

	for i := n - 1; i >= 0; i-- {
		for stack.Ok() && nums[stack.Peek()] > nums[i] {
			left[stack.Pop()] = i
		}
		stack.Push(i)
	}

	stack.Clear()

	right := make([]int, n)
	for i := range right {
		right[i] = n
	}

	for i := 0; i < n; i++ {
		for stack.Ok() && nums[stack.Peek()] > nums[i] {
			right[stack.Pop()] = i
		}
		stack.Push(i)
	}

	ans := 0
	for i := 0; i < n; i++ {
		if left[i] < k && k < right[i] {
			ans = max(ans, nums[i]*(right[i]-left[i]-1))
		}
	}

	return ans
}
