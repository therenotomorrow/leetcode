package golang

import (
	"reflect"
	"testing"
)

func TestRemoveNodes(t *testing.T) {
	t.Parallel()

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
				Val: 5,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 13,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  8,
								Next: nil,
							},
						},
					},
				},
			}},
			want: &ListNode{
				Val: 13,
				Next: &ListNode{
					Val:  8,
					Next: nil,
				},
			},
		},
		{
			name: "smoke 2",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			}},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := removeNodes(test.args.head)

			gotInts := make([]int, 0)
			for curr := got; curr != nil; curr = curr.Next {
				gotInts = append(gotInts, curr.Val)
			}

			wantInts := make([]int, 0)
			for curr := test.want; curr != nil; curr = curr.Next {
				wantInts = append(wantInts, curr.Val)
			}

			if !reflect.DeepEqual(gotInts, wantInts) {
				t.Errorf("removeNodes() = %v, want = %v", gotInts, wantInts)
			}
		})
	}
}
