package golang

import (
	"reflect"
	"testing"
)

func TestMergeNodes(t *testing.T) {
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
				Val: 0,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
									Next: &ListNode{
										Val: 2,
										Next: &ListNode{
											Val:  0,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
			}},
			want: &ListNode{Val: 4, Next: &ListNode{Val: 11, Next: nil}},
		},
		{
			name: "smoke 2",
			args: args{head: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 2,
									Next: &ListNode{
										Val: 2,
										Next: &ListNode{
											Val:  0,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
			}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := mergeNodes(test.args.head); !reflect.DeepEqual(got, test.want) {
				t.Errorf("mergeNodes() = %v, want = %v", got, test.want)
			}
		})
	}
}
