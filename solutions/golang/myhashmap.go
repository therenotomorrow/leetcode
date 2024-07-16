package golang

type MyHashMap struct {
	data []int
}

func MyHashMapConstructor() MyHashMap {
	const maxSize = 1_000_001

	data := make([]int, maxSize)
	for k := range data {
		data[k] = -1
	}

	return MyHashMap{data: data}
}

func (m *MyHashMap) Put(key int, value int) {
	m.data[key] = value
}

func (m *MyHashMap) Get(key int) int {
	return m.data[key]
}

func (m *MyHashMap) Remove(key int) {
	m.data[key] = -1
}
