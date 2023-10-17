package datastruct

type Set interface {
	Add(val int)
	Del(val int)
	Populate(vals ...int)
	Contains(val int) (ok bool)
}

type set struct {
	data map[int]struct{}
}

func NewSet() Set {
	return &set{data: make(map[int]struct{})}
}

func (s *set) Add(val int) {
	s.data[val] = struct{}{}
}

func (s *set) Del(val int) {
	delete(s.data, val)
}

func (s *set) Populate(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

func (s *set) Contains(val int) (ok bool) {
	_, ok = s.data[val]
	return
}
