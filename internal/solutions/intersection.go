package solutions

func intersection(nums1 []int, nums2 []int) []int {
	set1 := NewSet[int]()
	set2 := NewSet[int]()

	for _, num := range nums1 {
		set1.Add(num)
	}

	for _, num := range nums2 {
		if set1.Contains(num) {
			set2.Add(num)
		}
	}

	return set2.Values()
}

type Set[T comparable] interface {
	Add(val T)
	Del(val T)
	Size() int
	Populate(vals ...T)
	Contains(val T) bool
	Values() []T
}

type set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	return &set[T]{data: make(map[T]struct{})}
}

func (s *set[T]) Add(val T) {
	s.data[val] = struct{}{}
}

func (s *set[T]) Del(val T) {
	delete(s.data, val)
}

func (s *set[T]) Size() int {
	return len(s.data)
}

func (s *set[T]) Populate(vals ...T) {
	for _, val := range vals {
		s.Add(val)
	}
}

func (s *set[T]) Contains(val T) bool {
	_, ok := s.data[val]

	return ok
}

func (s *set[T]) Values() []T {
	vals := make([]T, 0, s.Size())

	for val := range s.data {
		vals = append(vals, val)
	}

	return vals
}
