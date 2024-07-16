package golang

type MyHashSet struct {
	data []int
}

func MyHashSetConstructor() MyHashSet {
	const maxSize = 1_000_001

	data := make([]int, maxSize)
	for i := range data {
		data[i] = -1
	}

	return MyHashSet{data: data}
}

func (s *MyHashSet) Add(key int) {
	s.data[key] = key
}

func (s *MyHashSet) Remove(key int) {
	s.data[key] = -1
}

func (s *MyHashSet) Contains(key int) bool {
	return s.data[key] != -1
}
