package golang

import "testing"

func TestGameResult(t *testing.T) {
	t.Parallel()

	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "smoke 1",
			args: args{head: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}},
			want: "Even",
		},
		{
			name: "smoke 2",
			args: args{head: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 7,
							Next: &ListNode{
								Val:  20,
								Next: &ListNode{Val: 5, Next: nil},
							},
						},
					},
				},
			}},
			want: "Odd",
		},
		{
			name: "smoke 3",
			args: args{head: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{Val: 1, Next: nil},
					},
				},
			}},
			want: "Tie",
		},
		{
			name: "test 439: wrong answer",
			args: args{head: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 27,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{Val: 35, Next: nil},
					},
				},
			}},
			want: "Odd",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := gameResult(test.args.head); got != test.want {
				t.Errorf("gameResult() = %v, want = %v", got, test.want)
			}
		})
	}
}
