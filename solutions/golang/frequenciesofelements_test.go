package golang

import (
	"reflect"
	"sort"
	"testing"
)

func TestFrequenciesOfElements(t *testing.T) {
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
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val:  2,
								Next: &ListNode{Val: 3, Next: nil},
							},
						},
					},
				},
			}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}},
		},
		{
			name: "smoke 2",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  2,
							Next: &ListNode{Val: 2, Next: nil},
						},
					},
				},
			}},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}},
		},
		{
			name: "smoke 3",
			args: args{head: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  2,
								Next: &ListNode{Val: 1, Next: nil},
							},
						},
					},
				},
			}},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val:  1,
								Next: &ListNode{Val: 1, Next: nil},
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

			head := frequenciesOfElements(test.args.head)
			got := make([]int, 0)

			for ; head != nil; head = head.Next {
				got = append(got, head.Val)
			}

			head = test.want
			want := make([]int, 0)

			for ; head != nil; head = head.Next {
				want = append(want, head.Val)
			}

			sort.Ints(got)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("frequenciesOfElements() = %v, want = %v", got, want)
			}
		})
	}
}
