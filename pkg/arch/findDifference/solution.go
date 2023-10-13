package findDifference

type Set struct {
	data map[int]struct{}
}

func NewSet(values ...int) *Set {
	s := Set{data: make(map[int]struct{})}

	for _, val := range values {
		s.data[val] = struct{}{}
	}

	return &s
}

func (s *Set) List() []int {
	i := 0
	values := make([]int, len(s.data))

	for val := range s.data {
		values[i] = val
		i++
	}

	return values
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	s1 := NewSet(nums1...)
	s2 := NewSet(nums2...)

	for val := range s1.data {
		if _, ok := s2.data[val]; ok {
			delete(s1.data, val)
			delete(s2.data, val)
		}
	}

	return [][]int{s1.List(), s2.List()}
}
