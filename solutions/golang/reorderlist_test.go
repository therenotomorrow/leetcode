package golang

import (
	"reflect"
	"testing"
)

func TestReorderList(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := make([]int, 0)

			reorderList(test.args.head)

			for curr := test.args.head; curr != nil; curr = curr.Next {
				got = append(got, curr.Val)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("reorderList() = %v, want = %v", got, test.want)
			}
		})
	}
}
