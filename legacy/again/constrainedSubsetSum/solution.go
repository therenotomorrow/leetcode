package constrainedSubsetSum

type dNode struct {
	val  int
	prev *dNode
	next *dNode
}

type deque struct {
	head *dNode
	tail *dNode
	size int
}

func (q *deque) Append(val int) {
	old := q.tail
	q.tail = &dNode{val: val}

	if q.Size() == 0 {
		q.head = q.tail
	} else {
		q.tail.prev = old
		old.next = q.tail
	}

	q.size++
}

func (q *deque) PopLeft() (val int, ok bool) {
	if q.Size() == 0 {
		return
	}

	val = q.head.val
	q.head = q.head.next

	if q.head != nil {
		q.head.prev = nil
	}

	q.size--

	return val, true
}

func (q *deque) PopRight() (val int, ok bool) {
	if q.Size() == 0 {
		return
	}

	val = q.tail.val
	q.tail = q.tail.prev

	if q.tail != nil {
		q.tail.next = nil
	}

	q.size--

	return val, true
}

func (q *deque) Peek() (val int, ok bool) {
	if q.Size() == 0 {
		return
	}

	return q.head.val, true
}

func (q *deque) PeekMust() (val int) {
	return q.head.val
}

func (q *deque) PopMust() (val int) {
	return q.tail.val
}

func (q *deque) Size() int {
	return q.size
}

func (q *deque) IsEmpty() bool {
	return q.Size() == 0
}

func constrainedSubsetSum(nums []int, k int) int {
	q := deque{}
	dynamic := make([]int, len(nums))

	for i, num := range nums {
		if !q.IsEmpty() && i-q.PeekMust() > k {
			q.PopLeft()
		}

		if !q.IsEmpty() {
			dynamic[i] = dynamic[q.PeekMust()]
		} else {
			dynamic[i] = 0
		}

		dynamic[i] += num

		for !q.IsEmpty() && dynamic[q.PopMust()] < dynamic[i] {
			q.PopRight()
		}

		if dynamic[i] > 0 {
			q.Append(i)
		}
	}

	m := dynamic[0]
	for _, d := range dynamic {
		if d > m {
			m = d
		}
	}

	return m
}
