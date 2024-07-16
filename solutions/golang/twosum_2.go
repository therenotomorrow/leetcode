package golang

type TwoSum struct {
	data map[int]int
}

func TwoSumConstructor() TwoSum {
	return TwoSum{data: make(map[int]int)}
}

func (s *TwoSum) Add(number int) {
	s.data[number]++
}

func (s *TwoSum) Find(value int) bool {
	for num := range s.data {
		cnt, exist := s.data[value-num]

		if value == 2*num {
			if cnt > 1 {
				return true
			}

			continue
		}

		if exist {
			return true
		}
	}

	return false
}
