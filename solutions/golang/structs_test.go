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
