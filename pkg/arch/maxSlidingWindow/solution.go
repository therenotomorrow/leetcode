package maxSlidingWindow

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{data: make([]int, 0)}
}

func (q *Queue) First() int {
	return q.data[0]
}

func (q *Queue) Last() int {
	return q.data[q.Len()-1]
}

func (q *Queue) Enqueue(val int) {
	q.data = append(q.data, val)
}

func (q *Queue) Dequeue() int {
	val := q.First()
	q.data = q.data[1:]

	return val
}

func (q *Queue) Pop() int {
	val := q.Last()
	q.data = q.data[:q.Len()-1]

	return val
}

func (q *Queue) Len() int {
	return len(q.data)
}

func (q *Queue) Next() bool {
	return q.Len() > 0
}

func maxSlidingWindow(nums []int, k int) []int {
	q := NewQueue()
	ans := make([]int, 0)

	for i, num := range nums {
		for q.Next() && num > nums[q.Last()] {
			q.Pop()
		}

		q.Enqueue(i)

		if q.First()+k == i {
			q.Dequeue()
		}

		if i >= k-1 {
			ans = append(ans, nums[q.First()])
		}
	}

	return ans
}
