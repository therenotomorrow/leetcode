package golang

type CustomStack struct {
	data    []int
	maxSize int
}

func CustomStackConstructor(maxSize int) CustomStack {
	return CustomStack{data: make([]int, 0, maxSize), maxSize: maxSize}
}

func (cs *CustomStack) Push(x int) {
	if len(cs.data) == cs.maxSize {
		return
	}

	cs.data = append(cs.data, x)
}

func (cs *CustomStack) Pop() int {
	if len(cs.data) == 0 {
		return -1
	}

	last := len(cs.data) - 1
	val := cs.data[last]
	cs.data = cs.data[:last]

	return val
}

func (cs *CustomStack) Increment(k int, val int) {
	i := 0

	for i < len(cs.data) && k > 0 {
		cs.data[i] += val
		i++
		k--
	}
}
