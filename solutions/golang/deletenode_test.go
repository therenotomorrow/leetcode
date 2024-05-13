package golang

import (
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
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
			name: "smoke 1",
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
			name: "smoke 2",
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

	for _, tt := range tests {
		switch tt.name {
		case "smoke 1":
			tt.list.Next = tt.args.node

		case "smoke 2":
			tt.list.Next.Next = tt.args.node
		}

		t.Run(tt.name, func(t *testing.T) {
			deleteNode(tt.args.node)

			if !reflect.DeepEqual(tt.list, tt.want) {
				t.Errorf("deleteNode() = %v, want = %v", tt.list, tt.want)
			}
		})
	}
}
