package golang

import (
	"crypto/rand"
	"math/big"
)

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
	idx, exist := rs.indexes[val]
	if !exist {
		return false
	}

	lastVal := rs.data[len(rs.data)-1]
	rs.data[idx] = lastVal
	rs.indexes[lastVal] = idx

	rs.data = rs.data[:len(rs.data)-1]
	delete(rs.indexes, val)

	return exist
}

func (rs *RandomizedSet) GetRandom() int {
	idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(rs.data))))

	return rs.data[int(idx.Int64())]
}
