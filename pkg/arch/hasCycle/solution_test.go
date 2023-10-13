package hasCycle

import (
	"testing"
)

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{head: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: -4,
						},
					},
				},
			}},
			want: true,
		},
		{
			name: "2",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			}},
			want: true,
		},
		{
			name: "3",
			args: args{head: &ListNode{
				Val: 1,
			}},
			want: false,
		},
		{
			name: "4",
			args: args{head: nil},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "1":
				tt.args.head.Next.Next.Next.Next = tt.args.head.Next

			case "2":
				tt.args.head.Next = tt.args.head
			}

			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
