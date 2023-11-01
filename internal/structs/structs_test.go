package structs_test

import (
	"github.com/therenotomorrow/leetcode/internal/structs"
	"testing"
)

func TestListNode(t *testing.T) {
	// make sure about creation

	_ = structs.ListNode{Val: 42, Next: &structs.ListNode{Val: 36}}
}

func TestTreeNode(t *testing.T) {
	// make sure about creation

	_ = structs.TreeNode{Val: 42, Left: &structs.TreeNode{Val: 36}, Right: &structs.TreeNode{Val: 25}}
}
