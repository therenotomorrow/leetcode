package golang

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type PairNode [2]int

type Node struct {
	Val      int
	Children []*Node
}

type DoubleListNode struct {
	Val  int
	Next *DoubleListNode
	Prev *DoubleListNode
}
