package NestedIterator

type NestedInteger struct{}

// IsInteger returns `true` if `nList` holds a single integer, rather than a nested list.
func (nList *NestedInteger) IsInteger() bool {
	return false
}

// GetInteger returns the single integer that `nList holds, if it holds a single integer.
// The result is undefined if `nList` holds a nested list.
func (nList *NestedInteger) GetInteger() int {
	return 0
}

// SetInteger sets `nList` to hold a single integer.
func (nList *NestedInteger) SetInteger(value int) {}

// Add adds `nList` to hold a nested list and adds a nested integer to it.
func (nList *NestedInteger) Add(elem NestedInteger) {}

// GetList returns the nested list that `nList` holds, if it holds a nested list.
// The list length is zero if this NestedInteger holds a single integer.
func (nList *NestedInteger) GetList() []*NestedInteger {
	return nil
}

type NestedIterator struct {
	data []int
	pos  int
}

func flatten(nList []*NestedInteger) []int {
	arr := make([]int, 0)

	for _, elem := range nList {
		if elem.IsInteger() {
			arr = append(arr, elem.GetInteger())
		} else {
			arr = append(arr, flatten(elem.GetList())...)
		}
	}

	return arr
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{data: flatten(nestedList), pos: -1}
}

func (iter *NestedIterator) Next() int {
	iter.pos++
	return iter.data[iter.pos]
}

func (iter *NestedIterator) HasNext() bool {
	return iter.pos+1 < len(iter.data)
}

// fixme: need normal tests
