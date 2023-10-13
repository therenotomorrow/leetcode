package LRUCache

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

type LRUCache struct {
	cap  int
	head *Node
	tail *Node
	data map[int]*Node
}

func Constructor(cap int) LRUCache {
	c := LRUCache{
		cap:  cap,
		head: &Node{Key: -1, Val: -1},
		tail: &Node{Key: -2, Val: -2},
		data: make(map[int]*Node),
	}

	c.head.Next = c.tail
	c.tail.Prev = c.head

	return c
}

func (c *LRUCache) add(node *Node) {
	prevTail := c.tail.Prev
	prevTail.Next = node

	node.Prev = prevTail
	node.Next = c.tail

	c.tail.Prev = node
}

func (c *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (c *LRUCache) Get(key int) int {
	node := c.data[key]

	if node == nil {
		return -1
	}

	c.remove(node)
	c.add(node)

	return node.Val
}

func (c *LRUCache) Put(key int, val int) {
	node := c.data[key]

	if node != nil {
		c.remove(node)
	}

	node = &Node{Key: key, Val: val}
	c.data[key] = node
	c.add(node)

	if len(c.data) > c.cap {
		node = c.head.Next
		c.remove(node)
		delete(c.data, node.Key)
	}
}
