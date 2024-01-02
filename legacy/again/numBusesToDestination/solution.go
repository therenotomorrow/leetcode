package numBusesToDestination

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	graph := make(map[int][]int)

	for r := 0; r < len(routes); r++ {
		for _, stop := range routes[r] {
			route, ok := graph[stop]
			if !ok {
				route = make([]int, 0)
			}
			route = append(route, r)
			graph[stop] = route
		}
	}

	q := NewQueue()
	vis := NewSet()

	for _, route := range graph[source] {
		q.Enqueue(route)
		vis.Add(route)
	}

	busCount := 1

	for !q.IsEmpty() {
		size := q.Size()

		for i := 0; i < size; i++ {
			route, _ := q.Dequeue()

			for _, stop := range routes[route] {
				if stop == target {
					return busCount
				}

				for _, next := range graph[stop] {
					if vis.Contains(next) {
						continue
					}

					vis.Add(next)
					q.Enqueue(next)
				}
			}
		}

		busCount++
	}

	return -1
}

type Queue interface {
	Enqueue(val int)
	Dequeue() (val int, ok bool)
	Peek() (val int, ok bool)
	Size() int
	IsEmpty() bool
}

type qNode struct {
	val  int
	next *qNode
}

type queue struct {
	head *qNode
	tail *qNode
	size int
}

func NewQueue() Queue {
	return &queue{}
}

func (q *queue) Enqueue(val int) {
	old := q.tail
	q.tail = &qNode{val: val}

	if q.IsEmpty() {
		q.head = q.tail
	} else {
		old.next = q.tail
	}

	q.size++
}

func (q *queue) Dequeue() (val int, ok bool) {
	if q.IsEmpty() {
		return
	}

	val = q.head.val
	q.head = q.head.next
	q.size--

	return val, true
}

func (q *queue) Peek() (val int, ok bool) {
	if q.IsEmpty() {
		return
	}

	return q.head.val, true
}

func (q *queue) Size() int {
	return q.size
}

func (q *queue) IsEmpty() bool {
	return q.Size() == 0
}

type Set interface {
	Add(val int)
	Del(val int)
	Populate(vals ...int)
	Contains(val int) (ok bool)
}

type set struct {
	data map[int]struct{}
}

func NewSet() Set {
	return &set{data: make(map[int]struct{})}
}

func (s *set) Add(val int) {
	s.data[val] = struct{}{}
}

func (s *set) Del(val int) {
	delete(s.data, val)
}

func (s *set) Populate(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

func (s *set) Contains(val int) (ok bool) {
	_, ok = s.data[val]
	return
}
