package golang

import (
	"testing"
)

func TestLeafSimilar(t *testing.T) {
	t.Parallel()

	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "smoke 1",
			args: args{
				root1: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:  5,
						Left: &TreeNode{Val: 6, Left: nil, Right: nil},
						Right: &TreeNode{
							Val:   2,
							Left:  &TreeNode{Val: 7, Left: nil, Right: nil},
							Right: &TreeNode{Val: 4, Left: nil, Right: nil},
						},
					},
					Right: &TreeNode{
						Val:   1,
						Left:  &TreeNode{Val: 9, Left: nil, Right: nil},
						Right: &TreeNode{Val: 8, Left: nil, Right: nil},
					},
				},
				root2: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   5,
						Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
						Right: &TreeNode{Val: 7, Left: nil, Right: nil},
					},
					Right: &TreeNode{
						Val:  1,
						Left: &TreeNode{Val: 4, Left: nil, Right: nil},
						Right: &TreeNode{
							Val:   2,
							Left:  &TreeNode{Val: 9, Left: nil, Right: nil},
							Right: &TreeNode{Val: 8, Left: nil, Right: nil},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "smoke 2",
			args: args{
				root1: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
					Right: &TreeNode{Val: 3, Left: nil, Right: nil},
				},
				root2: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
					Right: &TreeNode{Val: 2, Left: nil, Right: nil},
				},
			},
			want: false,
		},
		{
			name: "test 27: wrong answer",
			args: args{
				root1: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   5,
						Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
						Right: &TreeNode{Val: 7, Left: nil, Right: nil},
					},
					Right: &TreeNode{
						Val:  1,
						Left: &TreeNode{Val: 4, Left: nil, Right: nil},
						Right: &TreeNode{
							Val:  2,
							Left: &TreeNode{Val: 9, Left: nil, Right: nil},
							Right: &TreeNode{
								Val:   11,
								Left:  &TreeNode{Val: 8, Left: nil, Right: nil},
								Right: &TreeNode{Val: 10, Left: nil, Right: nil},
							},
						},
					},
				},
				root2: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:  5,
						Left: &TreeNode{Val: 6, Left: nil, Right: nil},
						Right: &TreeNode{
							Val:   2,
							Left:  &TreeNode{Val: 7, Left: nil, Right: nil},
							Right: &TreeNode{Val: 4, Left: nil, Right: nil},
						},
					},
					Right: &TreeNode{
						Val:   1,
						Left:  &TreeNode{Val: 9, Left: nil, Right: nil},
						Right: &TreeNode{Val: 8, Left: nil, Right: nil},
					},
				},
			},
			want: false,
		},
		{
			name: "test 29: wrong answer",
			args: args{
				root1: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
						Right: nil,
					},
					Right: nil,
				},
				root2: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
					Right: nil,
				},
			},
			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := leafSimilar(test.args.root1, test.args.root2); got != test.want {
				t.Errorf("leafSimilar() = %v, want = %v", got, test.want)
			}
		})
	}
}
