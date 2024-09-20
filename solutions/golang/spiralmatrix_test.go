package golang

import (
	"reflect"
	"testing"
)

func TestSpiralMatrix(t *testing.T) {
	t.Parallel()

	type args struct {
		m    int
		n    int
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "smoke 1",
			args: args{m: 3, n: 5, head: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 8,
								Next: &ListNode{
									Val: 1,
									Next: &ListNode{
										Val: 7,
										Next: &ListNode{
											Val: 9,
											Next: &ListNode{
												Val: 4,
												Next: &ListNode{
													Val: 2,
													Next: &ListNode{
														Val: 5,
														Next: &ListNode{
															Val:  5,
															Next: &ListNode{Val: 0, Next: nil},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			}},
			want: [][]int{{3, 0, 2, 6, 8}, {5, 0, -1, -1, 1}, {5, 2, 4, 9, 7}},
		},
		{
			name: "smoke 2",
			args: args{m: 1, n: 4, head: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  1,
					Next: &ListNode{Val: 2, Next: nil},
				},
			}},
			want: [][]int{{0, 1, 2, -1}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := spiralMatrix(test.args.m, test.args.n, test.args.head); !reflect.DeepEqual(got, test.want) {
				t.Errorf("spiralMatrix() = %v, want = %v", got, test.want)
			}
		})
	}
}
