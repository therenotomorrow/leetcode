package reverseEvenLengthGroups

import (
	"reflect"
	"testing"
)

func Test_reverseEvenLengthGroups(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{head: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 1,
									Next: &ListNode{
										Val: 7,
										Next: &ListNode{
											Val: 3,
											Next: &ListNode{
												Val: 8,
												Next: &ListNode{
													Val: 4,
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
			want: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 1,
									Next: &ListNode{
										Val: 4,
										Next: &ListNode{
											Val: 8,
											Next: &ListNode{
												Val: 3,
												Next: &ListNode{
													Val: 7,
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
		{
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			}},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 6,
							},
						},
					},
				},
			},
		},
		{
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			}},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseEvenLengthGroups(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseEvenLengthGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
