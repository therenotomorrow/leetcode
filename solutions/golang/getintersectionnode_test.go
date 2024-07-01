package golang

import (
	"reflect"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	t.Parallel()

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
			name: Smoke1,
			args: args{
				headA: &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: nil}},
				headB: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: nil}}},
			},
			want: nil,
		},
		{
			name: Smoke2,
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

	for _, test := range tests {
		switch test.name {
		case Smoke1:
			inter := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}

			test.args.headA.Next.Next = inter
			test.args.headB.Next.Next.Next = inter

			test.want = inter

		case Smoke2:
			inter := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}

			test.args.headA.Next.Next.Next = inter
			test.args.headB.Next = inter

			test.want = inter
		}

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := getIntersectionNode(test.args.headA, test.args.headB); !reflect.DeepEqual(got, test.want) {
				t.Errorf("getIntersectionNode() = %v, want = %v", got, test.want)
			}
		})
	}
}
