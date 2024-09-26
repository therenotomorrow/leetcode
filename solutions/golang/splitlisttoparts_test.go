package golang

import (
	"reflect"
	"testing"
)

func TestSplitListToParts(t *testing.T) {
	t.Parallel()

	type args struct {
		head *ListNode
		k    int
	}

	tests := []struct {
		name string
		args args
		want []*ListNode
	}{
		{
			name: "smoke 1",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{Val: 3, Next: nil},
					},
				},
				k: 5,
			},
			want: []*ListNode{
				{Val: 1, Next: nil},
				{Val: 2, Next: nil},
				{Val: 3, Next: nil},
				nil,
				nil,
			},
		},
		{
			name: "smoke 2",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
									Next: &ListNode{
										Val: 6,
										Next: &ListNode{
											Val: 7,
											Next: &ListNode{
												Val: 8,
												Next: &ListNode{
													Val:  9,
													Next: &ListNode{Val: 10, Next: nil},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				k: 3,
			},
			want: []*ListNode{
				{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}},
				{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: nil}}},
				{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 10, Next: nil}}},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := splitListToParts(test.args.head, test.args.k); !reflect.DeepEqual(got, test.want) {
				t.Errorf("splitListToParts() = %v, want = %v", got, test.want)
			}
		})
	}
}
