package golang

import "testing"

func TestHasCycle(t *testing.T) {
	t.Parallel()

	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: Smoke1,
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
			name: Smoke2,
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

	for _, test := range tests {
		switch test.name {
		case Smoke1:
			test.args.head.Next.Next.Next.Next = test.args.head.Next

		case Smoke2:
			test.args.head.Next = test.args.head
		}

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := hasCycle(test.args.head); got != test.want {
				t.Errorf("hasCycle() = %v, want = %v", got, test.want)
			}
		})
	}
}
