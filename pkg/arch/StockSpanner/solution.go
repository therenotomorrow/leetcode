package StockSpanner

type span = [2]int

type stack struct {
	data []span
}

func (s *stack) Len() int {
	return len(s.data)
}

func (s *stack) Last() int {
	return s.Len() - 1
}

func (s *stack) Next() bool {
	return s.Len() > 0
}

func (s *stack) Peek() span {
	return s.data[s.Len()-1]
}

func (s *stack) Push(val span) {
	s.data = append(s.data, val)
}

func (s *stack) Pop() span {
	data := s.data[s.Last()]
	s.data = s.data[:s.Last()]
	return data
}

type StockSpanner struct {
	s *stack
}

func Constructor() StockSpanner {
	return StockSpanner{s: &stack{}}
}

func (ss *StockSpanner) Next(price int) int {
	days := 1

	for ss.s.Next() && ss.s.Peek()[0] <= price {
		days += ss.s.Pop()[1]
	}
	ss.s.Push(span{price, days})

	return days
}
