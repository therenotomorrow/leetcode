package maxLevelSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	tn   *TreeNode
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
	len  int
}

func NewQueue(val *TreeNode) Queue {
	q := Queue{}
	q.Put(val)
	return q
}

func (q *Queue) Put(val *TreeNode) {
	old := q.tail
	q.tail = &Node{tn: val}

	if q.head == nil {
		q.head = q.tail
	} else {
		old.next = q.tail
	}

	q.len++
}

func (q *Queue) Get() *Node {
	val := q.head

	if val == nil {
		return val
	}

	q.head = q.head.next
	q.len--

	return val
}

func (q *Queue) IsFull() bool {
	return q.head != nil
}

func (q *Queue) Len() int {
	return q.len
}

func maxLevelSum(root *TreeNode) int {
	queue := NewQueue(root)

	max := -1_000_000
	lev := 0
	res := 0

	for queue.IsFull() {
		sum := 0
		lev += 1

		for i := queue.Len(); i > 0; i-- {
			val := queue.Get()

			sum += val.tn.Val

			if val.tn.Left != nil {
				queue.Put(val.tn.Left)
			}
			if val.tn.Right != nil {
				queue.Put(val.tn.Right)
			}
		}

		if sum > max {
			max = sum
			res = lev
		}
	}

	return res
}
