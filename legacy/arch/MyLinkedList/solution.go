package MyLinkedList

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type MyLinkedList struct {
	Head *Node
	Tail *Node
	Len  int
}

func Constructor() MyLinkedList {
	head := &Node{Val: -1}
	tail := &Node{Val: -2}

	head.Next = tail
	tail.Prev = head

	return MyLinkedList{Head: head, Tail: tail}
}

func (l *MyLinkedList) findNode(index int) *Node {
	curr := l.Head.Next
	for i := 0; i < index; i++ {
		curr = curr.Next
	}

	return curr
}

func (l *MyLinkedList) guard(index int) bool {
	return index >= l.Len
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.Len, val)
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if l.guard(index - 1) {
		return
	}

	node := l.findNode(index)

	forInsert := &Node{Val: val, Prev: node.Prev, Next: node}

	node.Prev.Next, node.Prev = forInsert, forInsert

	l.Len++
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if l.guard(index) {
		return
	}

	forDelete := l.findNode(index)

	forDelete.Prev.Next, forDelete.Next.Prev = forDelete.Next, forDelete.Prev

	l.Len--
}

func (l *MyLinkedList) Get(index int) int {
	if l.guard(index) {
		return -1
	}

	return l.findNode(index).Val
}
