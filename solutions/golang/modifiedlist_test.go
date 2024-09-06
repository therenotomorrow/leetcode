package golang

import (
	"reflect"
	"testing"
)

func TestModifiedList(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "smoke 1",
			args: args{
				nums: []int{1, 2, 3},
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  4,
								Next: &ListNode{Val: 5, Next: nil},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val:  4,
				Next: &ListNode{Val: 5, Next: nil},
			},
		},
		{
			name: "smoke 2",
			args: args{
				nums: []int{1},
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 2,
								Next: &ListNode{
									Val:  1,
									Next: &ListNode{Val: 2, Next: nil},
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  2,
					Next: &ListNode{Val: 2, Next: nil},
				},
			},
		},
		{
			name: "smoke 3",
			args: args{
				nums: []int{5},
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  3,
							Next: &ListNode{Val: 4, Next: nil},
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: &ListNode{Val: 4, Next: nil},
					},
				},
			},
		},
		{
			name: "test 435: wrong answer",
			args: args{
				nums: []int{9, 2, 5},
				head: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  10,
						Next: &ListNode{Val: 9, Next: nil},
					},
				},
			},
			want: &ListNode{Val: 10, Next: nil},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := modifiedList(test.args.nums, test.args.head); !reflect.DeepEqual(got, test.want) {
				t.Errorf("modifiedList() = %v, want = %v", got, test.want)
			}
		})
	}
}
