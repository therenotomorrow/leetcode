package golang

import "testing"

func TestIsSubPath(t *testing.T) {
	t.Parallel()

	type args struct {
		head *ListNode
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "smoke 1",
			args: args{
				head: &ListNode{
					Val: 4, Next: &ListNode{
						Val: 2, Next: &ListNode{
							Val: 8, Next: nil,
						},
					},
				},
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  4,
						Left: nil,
						Right: &TreeNode{
							Val:   2,
							Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:  2,
							Left: &TreeNode{Val: 6, Left: nil, Right: nil},
							Right: &TreeNode{
								Val:   8,
								Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
								Right: &TreeNode{Val: 3, Left: nil, Right: nil},
							},
						},
						Right: nil,
					},
				},
			},
			want: true,
		},
		{
			name: "smoke 2",
			args: args{
				head: &ListNode{
					Val: 1, Next: &ListNode{
						Val: 4, Next: &ListNode{
							Val: 2, Next: &ListNode{
								Val: 6, Next: nil,
							},
						},
					},
				},
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  4,
						Left: nil,
						Right: &TreeNode{
							Val:   2,
							Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:  2,
							Left: &TreeNode{Val: 6, Left: nil, Right: nil},
							Right: &TreeNode{
								Val:   8,
								Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
								Right: &TreeNode{Val: 3, Left: nil, Right: nil},
							},
						},
						Right: nil,
					},
				},
			},
			want: true,
		},
		{
			name: "smoke 3",
			args: args{
				head: &ListNode{
					Val: 1, Next: &ListNode{
						Val: 4, Next: &ListNode{
							Val: 2, Next: &ListNode{
								Val: 6, Next: &ListNode{
									Val: 8, Next: nil,
								},
							},
						},
					},
				},
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  4,
						Left: nil,
						Right: &TreeNode{
							Val:   2,
							Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:  2,
							Left: &TreeNode{Val: 6, Left: nil, Right: nil},
							Right: &TreeNode{
								Val:   8,
								Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
								Right: &TreeNode{Val: 3, Left: nil, Right: nil},
							},
						},
						Right: nil,
					},
				},
			},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isSubPath(test.args.head, test.args.root); got != test.want {
				t.Errorf("isSubPath() = %v, want = %v", got, test.want)
			}
		})
	}
}
