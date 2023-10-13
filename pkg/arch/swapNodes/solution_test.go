package swapNodes

import (
	"reflect"
	"testing"
)

func Test_swapNodes(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
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
			}, k: 2},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
		{
			args: args{head: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 7,
								Next: &ListNode{
									Val: 8,
									Next: &ListNode{
										Val: 3,
										Next: &ListNode{
											Val: 0,
											Next: &ListNode{
												Val: 9,
												Next: &ListNode{
													Val: 5,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			}, k: 5},
			want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 8,
								Next: &ListNode{
									Val: 7,
									Next: &ListNode{
										Val: 3,
										Next: &ListNode{
											Val: 0,
											Next: &ListNode{
												Val: 9,
												Next: &ListNode{
													Val: 5,
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapNodes(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
