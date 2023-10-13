package fastAndSlowPointers

import "testing"

func Test_getMiddle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{&ListNode{
			val: 1,
			next: &ListNode{
				val: 2,
				next: &ListNode{
					val: 3,
					next: &ListNode{
						val: 4,
						next: &ListNode{
							val: 5,
						},
					},
				},
			},
		}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMiddle(tt.args.head); got != tt.want {
				t.Errorf("getMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findNode(t *testing.T) {
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
				val: 1,
				next: &ListNode{
					val: 2,
					next: &ListNode{
						val: 3,
						next: &ListNode{
							val: 4,
							next: &ListNode{
								val: 5,
							},
						},
					},
				},
			}, k: 2},
			want: &ListNode{val: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNode(tt.args.head, tt.args.k); got.val != tt.want.val {
				t.Errorf("findNode() = %v, want %v", got.val, tt.want.val)
			}
		})
	}
}
