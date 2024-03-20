package golang

import (
	"reflect"
	"testing"
)

func TestMergeInBetween(t *testing.T) {
	type args struct {
		list1 *ListNode
		a     int
		b     int
		list2 *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "smoke 1",
			args: args{
				list1: &ListNode{
					Val: 10, Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 13,
							Next: &ListNode{
								Val: 6,
								Next: &ListNode{
									Val:  9,
									Next: &ListNode{Val: 5, Next: nil},
								},
							},
						},
					},
				},
				a: 3,
				b: 4,
				list2: &ListNode{
					Val: 1000000,
					Next: &ListNode{
						Val:  1000001,
						Next: &ListNode{Val: 1000002, Next: nil},
					},
				},
			},
			want: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 13,
						Next: &ListNode{
							Val: 1000000,
							Next: &ListNode{
								Val: 1000001,
								Next: &ListNode{
									Val:  1000002,
									Next: &ListNode{Val: 5, Next: nil},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "smoke 2",
			args: args{
				list1: &ListNode{
					Val: 0, Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val:  5,
										Next: &ListNode{Val: 6, Next: nil},
									},
								},
							},
						},
					},
				},
				a: 2,
				b: 5,
				list2: &ListNode{
					Val: 1000000,
					Next: &ListNode{
						Val: 1000001,
						Next: &ListNode{
							Val: 1000002,
							Next: &ListNode{
								Val:  1000003,
								Next: &ListNode{Val: 1000004, Next: nil},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1000000,
						Next: &ListNode{
							Val: 1000001,
							Next: &ListNode{
								Val: 1000002,
								Next: &ListNode{
									Val: 1000003,
									Next: &ListNode{
										Val:  1000004,
										Next: &ListNode{Val: 6, Next: nil},
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
			if got := mergeInBetween(tt.args.list1, tt.args.a, tt.args.b, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeInBetween() = %v, want = %v", got, tt.want)
			}
		})
	}
}
