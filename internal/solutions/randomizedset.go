package solutions

import "math/rand"

type RandomizedSet struct {
	data    []int
	indexes map[int]int
}

func RandomizedSetConstructor() RandomizedSet {
	return RandomizedSet{data: make([]int, 0), indexes: make(map[int]int)}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.indexes[val]; ok {
		return false
	}

	rs.indexes[val] = len(rs.data)
	rs.data = append(rs.data, val)

	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	idx, ok := rs.indexes[val]
	if !ok {
		return false
	}

	lastVal := rs.data[len(rs.data)-1]
	rs.data[idx] = lastVal
	rs.indexes[lastVal] = idx

	rs.data = rs.data[:len(rs.data)-1]
	delete(rs.indexes, val)

	return ok
}

func (rs *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(rs.data))

	return rs.data[idx]
}
