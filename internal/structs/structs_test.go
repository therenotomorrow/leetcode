package structs_test

import (
	"github.com/therenotomorrow/leetcode/internal/structs"
	"testing"
)

func TestListNode(t *testing.T) {
	// make sure about creation

	_ = structs.ListNode{Val: 42, Next: &structs.ListNode{Val: 36}}
}
