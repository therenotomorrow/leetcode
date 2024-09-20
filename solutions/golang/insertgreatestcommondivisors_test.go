package golang

import (
	"reflect"
	"testing"
)

func TestInsertGreatestCommonDivisors(t *testing.T) {
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
				Val: 18,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  10,
						Next: &ListNode{Val: 3, Next: nil},
					},
				},
			}},
			want: &ListNode{
				Val: 18,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 10,
								Next: &ListNode{
									Val:  1,
									Next: &ListNode{Val: 3, Next: nil},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "smoke 2",
			args: args{head: &ListNode{Val: 7, Next: nil}},
			want: &ListNode{Val: 7, Next: nil},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := insertGreatestCommonDivisors(test.args.head); !reflect.DeepEqual(got, test.want) {
				t.Errorf("insertGreatestCommonDivisors() = %v, want = %v", got, test.want)
			}
		})
	}
}
