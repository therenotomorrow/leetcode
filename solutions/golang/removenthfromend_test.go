package golang

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
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
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			}, n: 2},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		},
		{
			name: "smoke 2",
			args: args{head: &ListNode{Val: 1}, n: 1},
			want: nil,
		},
		{
			name: "smoke 3",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			}, n: 1},
			want: &ListNode{
				Val: 1,
			},
		},
		{
			name: "test 190: wrong answer",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			}, n: 2},
			want: &ListNode{
				Val: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want = %v", got, tt.want)
			}
		})
	}
}
