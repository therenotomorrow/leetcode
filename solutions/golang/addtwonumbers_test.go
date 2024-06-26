package golang

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Parallel()

	type args struct {
		l1 *ListNode
		l2 *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "smoke 1",
			args: args{
				l1: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{Val: 3, Next: nil},
					},
				},
				l2: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val:  6,
						Next: &ListNode{Val: 4, Next: nil},
					},
				},
			},
			want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val:  0,
					Next: &ListNode{Val: 8, Next: nil},
				},
			},
		},
		{
			name: "smoke 2",
			args: args{
				l1: &ListNode{Val: 0, Next: nil},
				l2: &ListNode{Val: 0, Next: nil},
			},
			want: &ListNode{Val: 0, Next: nil},
		},
		{
			name: "smoke 3",
			args: args{
				l1: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val:  9,
										Next: &ListNode{Val: 9, Next: nil},
									},
								},
							},
						},
					},
				},
				l2: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val:  9,
							Next: &ListNode{Val: 9, Next: nil},
						},
					},
				},
			},
			want: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9, Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val:  0,
										Next: &ListNode{Val: 1, Next: nil},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := addTwoNumbers(test.args.l1, test.args.l2); !reflect.DeepEqual(got, test.want) {
				t.Errorf("addTwoNumbers() = %v, want = %v", got, test.want)
			}
		})
	}
}
