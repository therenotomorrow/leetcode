package golang

import (
	"reflect"
	"testing"
)

func TestReorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smoke 1",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: &ListNode{Val: 4, Next: nil},
					},
				},
			}},
			want: []int{1, 4, 2, 3},
		},
		{
			name: "smoke 2",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: &ListNode{Val: 5, Next: nil},
						},
					},
				},
			}},
			want: []int{1, 5, 2, 4, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int, 0)

			reorderList(tt.args.head)

			for curr := tt.args.head; curr != nil; curr = curr.Next {
				got = append(got, curr.Val)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorderList() = %v, want = %v", got, tt.want)
			}
		})
	}
}
