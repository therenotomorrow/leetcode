package golang

import (
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	t.Parallel()

	type args struct {
		node *ListNode
	}

	tests := []struct {
		name string
		args args
		list *ListNode
		want *ListNode
	}{
		{
			name: Smoke1,
			args: args{node: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val:  1,
					Next: &ListNode{Val: 9, Next: nil},
				},
			}},
			list: &ListNode{Val: 4, Next: nil},
			want: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  1,
					Next: &ListNode{Val: 9, Next: nil},
				},
			},
		},
		{
			name: Smoke2,
			args: args{node: &ListNode{
				Val:  1,
				Next: &ListNode{Val: 9, Next: nil},
			}},
			list: &ListNode{
				Val:  4,
				Next: &ListNode{Val: 5, Next: nil},
			},
			want: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: &ListNode{Val: 9, Next: nil},
				},
			},
		},
	}

	for _, test := range tests {
		switch test.name {
		case Smoke1:
			test.list.Next = test.args.node

		case Smoke2:
			test.list.Next.Next = test.args.node
		}

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			deleteNode(test.args.node)

			if !reflect.DeepEqual(test.list, test.want) {
				t.Errorf("deleteNode() = %v, want = %v", test.list, test.want)
			}
		})
	}
}
