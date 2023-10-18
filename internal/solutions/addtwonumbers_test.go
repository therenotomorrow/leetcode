package solutions

import (
	"github.com/therenotomorrow/leetcode/internal/structs"
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	type args struct {
		l1 *structs.ListNode
		l2 *structs.ListNode
	}
	tests := []struct {
		name string
		args args
		want *structs.ListNode
	}{
		{
			name: "smoke 1",
			args: args{
				l1: &structs.ListNode{
					Val: 2,
					Next: &structs.ListNode{
						Val:  4,
						Next: &structs.ListNode{Val: 3},
					},
				},
				l2: &structs.ListNode{
					Val: 5,
					Next: &structs.ListNode{
						Val:  6,
						Next: &structs.ListNode{Val: 4},
					},
				},
			},
			want: &structs.ListNode{
				Val: 7,
				Next: &structs.ListNode{
					Val:  0,
					Next: &structs.ListNode{Val: 8},
				},
			},
		},
		{
			name: "smoke 2",
			args: args{
				l1: &structs.ListNode{Val: 0},
				l2: &structs.ListNode{Val: 0},
			},
			want: &structs.ListNode{Val: 0},
		},
		{
			name: "smoke 3",
			args: args{
				l1: &structs.ListNode{
					Val: 9,
					Next: &structs.ListNode{
						Val: 9,
						Next: &structs.ListNode{
							Val: 9,
							Next: &structs.ListNode{
								Val: 9,
								Next: &structs.ListNode{
									Val: 9,
									Next: &structs.ListNode{
										Val:  9,
										Next: &structs.ListNode{Val: 9},
									},
								},
							},
						},
					},
				},
				l2: &structs.ListNode{
					Val: 9,
					Next: &structs.ListNode{
						Val: 9,
						Next: &structs.ListNode{
							Val:  9,
							Next: &structs.ListNode{Val: 9},
						},
					},
				},
			},
			want: &structs.ListNode{
				Val: 8,
				Next: &structs.ListNode{
					Val: 9,
					Next: &structs.ListNode{
						Val: 9, Next: &structs.ListNode{
							Val: 9,
							Next: &structs.ListNode{
								Val: 0,
								Next: &structs.ListNode{
									Val: 0,
									Next: &structs.ListNode{
										Val:  0,
										Next: &structs.ListNode{Val: 1},
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
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want = %v", got, tt.want)
			}
		})
	}
}
