package datastruct

type Set[T int | rune] interface {
	Add(val T)
	Del(val T)
	Populate(vals ...T)
	Contains(val T) (ok bool)
}

type set[T int | rune] struct {
	data map[T]struct{}
}

func NewSet[T int | rune]() Set[T] {
	return &set[T]{data: make(map[T]struct{})}
}

func (s *set[T]) Add(val T) {
	s.data[val] = struct{}{}
}

func (s *set[T]) Del(val T) {
	delete(s.data, val)
}

func (s *set[T]) Populate(vals ...T) {
	for _, val := range vals {
		s.Add(val)
	}
}

func (s *set[T]) Contains(val T) (ok bool) {
	_, ok = s.data[val]
	return
}
