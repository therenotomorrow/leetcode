package golang

import "testing"

func TestIsPalindrome2(t *testing.T) {
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
			name: "smoke 1",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 1, Next: nil,
						},
					},
				},
			}},
			want: true,
		},
		{
			name: "smoke 2",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2, Next: nil,
				},
			}},
			want: false,
		},
		{
			name: "test 79: wrong answer",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 1, Next: nil,
						},
					},
				},
			}},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isPalindrome2(test.args.head); got != test.want {
				t.Errorf("isPalindrome2() = %v, want = %v", got, test.want)
			}
		})
	}
}
