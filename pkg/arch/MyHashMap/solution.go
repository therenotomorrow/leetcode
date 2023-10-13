package MyHashMap

type MyHashMap struct {
	data []int
}

const MapSize = 1_000_001

func Constructor() MyHashMap {
	data := make([]int, MapSize)

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
