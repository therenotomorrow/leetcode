package regionsBySlashes

func regionsBySlashes(grid []string) int {
	expanded := make([][]int, len(grid)*3)
	for i := range expanded {
		expanded[i] = make([]int, len(grid)*3)
	}

	filler := func(row int, col int) {
		dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

		expanded[row][col] = 1

		que := NewQueue[[]int]()
		que.Enqueue([]int{row, col})

		for !que.IsEmpty() {
			currentCell, _ := que.Dequeue()

			for _, direction := range dirs {
				i := direction[0] + currentCell[0]
				j := direction[1] + currentCell[1]

				if i >= 0 && j >= 0 && i < len(expanded) && j < len(expanded) && expanded[i][j] == 0 {
					expanded[i][j] = 1

					que.Enqueue([]int{i, j})
				}
			}
		}
	}

	for i := range len(grid) {
		for j := range len(grid) {
			switch row, col := i*3, j*3; grid[i][j] {
			case '\\':
				expanded[row][col] = 1
				expanded[row+1][col+1] = 1
				expanded[row+2][col+2] = 1
			case '/':
				expanded[row][col+2] = 1
				expanded[row+1][col+1] = 1
				expanded[row+2][col] = 1
			}
		}
	}

	cnt := 0

	for i := range len(expanded) {
		for j := range len(expanded) {
			if expanded[i][j] != 0 {
				continue
			}

			filler(i, j)

			cnt++
		}
	}

	return cnt
}

type Queueable interface {
	~[]int
}

type Queue[T Queueable] interface {
	Enqueue(val T)
	Dequeue() (T, bool)
	Peek() (T, bool)
	Size() int
	IsEmpty() bool
}

type qNode[T Queueable] struct {
	val  T
	next *qNode[T]
}

type queue[T Queueable] struct {
	head *qNode[T]
	tail *qNode[T]
	size int
}

func NewQueue[T Queueable]() Queue[T] {
	return &queue[T]{head: nil, tail: nil, size: 0}
}

func (q *queue[T]) Enqueue(val T) {
	old := q.tail
	q.tail = &qNode[T]{val: val, next: nil}

	if q.IsEmpty() {
		q.head = q.tail
	} else {
		old.next = q.tail
	}

	q.size++
}

func (q *queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T

		return zero, false
	}

	val := q.head.val
	q.head = q.head.next
	q.size--

	return val, true
}

func (q *queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var zero T

		return zero, false
	}

	return q.head.val, true
}

func (q *queue[T]) Size() int {
	return q.size
}

func (q *queue[T]) IsEmpty() bool {
	return q.Size() == 0
}
