package golang

import (
	"reflect"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "smoke 1",
			args: args{
				headA: &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: nil}},
				headB: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: nil}}},
			},
			want: nil,
		},
		{
			name: "smoke 2",
			args: args{
				headA: &ListNode{Val: 1, Next: &ListNode{Val: 9, Next: &ListNode{Val: 1, Next: nil}}},
				headB: &ListNode{Val: 3, Next: nil},
			},
			want: nil,
		},
		{
			name: "smoke 3",
			args: args{
				headA: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}},
				headB: &ListNode{Val: 1, Next: &ListNode{Val: 5, Next: nil}},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		switch tt.name {
		case "smoke 1":
			inter := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}

			tt.args.headA.Next.Next = inter
			tt.args.headB.Next.Next.Next = inter

			tt.want = inter

		case "smoke 2":
			inter := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}

			tt.args.headA.Next.Next.Next = inter
			tt.args.headB.Next = inter

			tt.want = inter
		}

		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want = %v", got, tt.want)
			}
		})
	}
}
