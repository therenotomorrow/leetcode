package dailyTemperatures

type Stack struct {
	data []int
}

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *Stack) Peek() int {
	if s.Len() == 0 {
		return 0
	}

	return s.data[s.Len()-1]
}

func (s *Stack) Pop() int {
	if s.Len() == 0 {
		return 0
	}

	val := s.data[s.Len()-1]
	s.data = s.data[:s.Len()-1]

	return val
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Have() bool {
	return s.Len() > 0
}

func dailyTemperatures(temperatures []int) []int {
	days := make([]int, len(temperatures))
	s := Stack{}

	for i, temp := range temperatures {
		for s.Have() && temperatures[s.Peek()] < temp {
			currDay := s.Pop()
			days[currDay] = i - currDay
		}

		s.Push(i)
	}

	return days
}
