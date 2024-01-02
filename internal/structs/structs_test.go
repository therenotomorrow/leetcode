package structs_test

import (
	"testing"

	"github.com/therenotomorrow/leetcode/internal/structs"
)

func TestListNode(_ *testing.T) {
	// make sure about creation
	_ = structs.ListNode{Val: 42, Next: &structs.ListNode{Val: 36, Next: nil}}
}

func TestTreeNode(_ *testing.T) {
	// make sure about creation
	_ = structs.TreeNode{
		Val:   42,
		Left:  &structs.TreeNode{Val: 36, Left: nil, Right: nil},
		Right: &structs.TreeNode{Val: 25, Left: nil, Right: nil},
	}
}
