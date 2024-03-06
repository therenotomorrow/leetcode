package golang

import "testing"

func TestHasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "smoke 1",
			args: args{head: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  0,
						Next: &ListNode{Val: -4, Next: nil},
					},
				},
			}},
			want: true,
		},
		{
			name: "smoke 2",
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}},
			want: true,
		},
		{
			name: "smoke 3",
			args: args{head: &ListNode{Val: 1, Next: nil}},
			want: false,
		},
		{
			name: "test 3: runtime",
			args: args{head: nil},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "smoke 1":
				tt.args.head.Next.Next.Next.Next = tt.args.head.Next

			case "smoke 2":
				tt.args.head.Next = tt.args.head
			}

			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want = %v", got, tt.want)
			}
		})
	}
}
