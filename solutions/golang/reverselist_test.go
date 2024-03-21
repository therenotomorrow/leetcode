package golang

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "smoke 1",
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
			want: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  2,
							Next: &ListNode{Val: 1, Next: nil},
						},
					},
				},
			},
		},
		{
			name: "smoke 2",
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}},
		},
		{
			name: "smoke 3",
			args: args{head: nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want = %v", got, tt.want)
			}
		})
	}
}
