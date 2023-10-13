package reversingALinkedList

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{args: args{head: &ListNode{
			val: 0,
			next: &ListNode{
				val: 1,
				next: &ListNode{
					val:  2,
					next: &ListNode{val: 3},
				},
			},
		}}, want: &ListNode{
			val: 3,
			next: &ListNode{
				val: 2,
				next: &ListNode{
					val:  1,
					next: &ListNode{val: 0},
				},
			},
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
