package golang_test

import (
	"testing"

	"github.com/therenotomorrow/leetcode/solutions/golang"
)

func TestListNode(t *testing.T) {
	t.Parallel()
	// make sure about creation
	_ = golang.ListNode{Val: 42, Next: &golang.ListNode{Val: 36, Next: nil}}
}

func TestTreeNode(t *testing.T) {
	t.Parallel()
	// make sure about creation
	_ = golang.TreeNode{
		Val:   42,
		Left:  &golang.TreeNode{Val: 36, Left: nil, Right: nil},
		Right: &golang.TreeNode{Val: 25, Left: nil, Right: nil},
	}
}

func TestPairNode(t *testing.T) {
	t.Parallel()
	// make sure about creation
	_ = golang.PairNode{1, 2}
}

func TestNode(t *testing.T) {
	t.Parallel()
	// make sure about creation
	_ = golang.Node{
		Val:      1,
		Children: []*golang.Node{{Val: 123, Children: nil}, {Val: 321, Children: nil}, {Val: 33, Children: nil}},
	}
}

func TestDoubleListNode(t *testing.T) {
	t.Parallel()
	// make sure about creation
	_ = golang.DoubleListNode{
		Val:  1,
		Next: &golang.DoubleListNode{Val: 2, Next: nil, Prev: nil},
		Prev: &golang.DoubleListNode{Val: 3, Next: nil, Prev: nil},
	}
}
