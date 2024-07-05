package golang

import (
	"reflect"
	"testing"
)

func TestNodesBetweenCriticalPoints(t *testing.T) {
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
				Val: 3,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			}},
			want: []int{-1, -1},
		},
		{
			name: "smoke 2",
			args: args{head: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val: 1,
									Next: &ListNode{
										Val:  2,
										Next: nil,
									},
								},
							},
						},
					},
				},
			}},
			want: []int{1, 3},
		},
		{
			name: "smoke 3",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 2,
									Next: &ListNode{
										Val: 2,
										Next: &ListNode{
											Val: 2,
											Next: &ListNode{
												Val:  7,
												Next: nil,
											},
										},
									},
								},
							},
						},
					},
				},
			}},
			want: []int{3, 3},
		},
		{
			name: "test 99: wrong answer",
			args: args{head: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 8,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 6,
									Next: &ListNode{
										Val: 6,
										Next: &ListNode{
											Val: 10,
											Next: &ListNode{
												Val:  6,
												Next: nil,
											},
										},
									},
								},
							},
						},
					},
				},
			}},
			want: []int{1, 6},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := nodesBetweenCriticalPoints(test.args.head); !reflect.DeepEqual(got, test.want) {
				t.Errorf("nodesBetweenCriticalPoints() = %v, want = %v", got, test.want)
			}
		})
	}
}
