package golang

import (
	"reflect"
	"testing"
)

func TestDoubleIt(t *testing.T) {
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
					Val: 8,
					Next: &ListNode{
						Val:  9,
						Next: nil,
					},
				},
			}},
			want: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
		{
			name: "smoke 2",
			args: args{head: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val:  9,
						Next: nil,
					},
				},
			}},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val:  8,
							Next: nil,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doubleIt(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doubleIt() = %v, want = %v", got, tt.want)
			}
		})
	}
}
