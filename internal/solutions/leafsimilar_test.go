package solutions

import (
	"testing"

	"github.com/therenotomorrow/leetcode/internal/structs"
)

func TestLeafSimilar(t *testing.T) {
	type args struct {
		root1 *structs.TreeNode
		root2 *structs.TreeNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "smoke 1",
			args: args{
				root1: &structs.TreeNode{
					Val: 3,
					Left: &structs.TreeNode{
						Val:  5,
						Left: &structs.TreeNode{Val: 6, Left: nil, Right: nil},
						Right: &structs.TreeNode{
							Val:   2,
							Left:  &structs.TreeNode{Val: 7, Left: nil, Right: nil},
							Right: &structs.TreeNode{Val: 4, Left: nil, Right: nil},
						},
					},
					Right: &structs.TreeNode{
						Val:   1,
						Left:  &structs.TreeNode{Val: 9, Left: nil, Right: nil},
						Right: &structs.TreeNode{Val: 8, Left: nil, Right: nil},
					},
				},
				root2: &structs.TreeNode{
					Val: 3,
					Left: &structs.TreeNode{
						Val:   5,
						Left:  &structs.TreeNode{Val: 6, Left: nil, Right: nil},
						Right: &structs.TreeNode{Val: 7, Left: nil, Right: nil},
					},
					Right: &structs.TreeNode{
						Val:  1,
						Left: &structs.TreeNode{Val: 4, Left: nil, Right: nil},
						Right: &structs.TreeNode{
							Val:   2,
							Left:  &structs.TreeNode{Val: 9, Left: nil, Right: nil},
							Right: &structs.TreeNode{Val: 8, Left: nil, Right: nil},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "smoke 2",
			args: args{
				root1: &structs.TreeNode{
					Val:   1,
					Left:  &structs.TreeNode{Val: 2, Left: nil, Right: nil},
					Right: &structs.TreeNode{Val: 3, Left: nil, Right: nil},
				},
				root2: &structs.TreeNode{
					Val:   1,
					Left:  &structs.TreeNode{Val: 3, Left: nil, Right: nil},
					Right: &structs.TreeNode{Val: 2, Left: nil, Right: nil},
				},
			},
			want: false,
		},
		{
			name: "test 27: wrong answer",
			args: args{
				root1: &structs.TreeNode{
					Val: 3,
					Left: &structs.TreeNode{
						Val:   5,
						Left:  &structs.TreeNode{Val: 6, Left: nil, Right: nil},
						Right: &structs.TreeNode{Val: 7, Left: nil, Right: nil},
					},
					Right: &structs.TreeNode{
						Val:  1,
						Left: &structs.TreeNode{Val: 4, Left: nil, Right: nil},
						Right: &structs.TreeNode{
							Val:  2,
							Left: &structs.TreeNode{Val: 9, Left: nil, Right: nil},
							Right: &structs.TreeNode{
								Val:   11,
								Left:  &structs.TreeNode{Val: 8, Left: nil, Right: nil},
								Right: &structs.TreeNode{Val: 10, Left: nil, Right: nil},
							},
						},
					},
				},
				root2: &structs.TreeNode{
					Val: 3,
					Left: &structs.TreeNode{
						Val:  5,
						Left: &structs.TreeNode{Val: 6, Left: nil, Right: nil},
						Right: &structs.TreeNode{
							Val:   2,
							Left:  &structs.TreeNode{Val: 7, Left: nil, Right: nil},
							Right: &structs.TreeNode{Val: 4, Left: nil, Right: nil},
						},
					},
					Right: &structs.TreeNode{
						Val:   1,
						Left:  &structs.TreeNode{Val: 9, Left: nil, Right: nil},
						Right: &structs.TreeNode{Val: 8, Left: nil, Right: nil},
					},
				},
			},
			want: false,
		},
		{
			name: "test 29: wrong answer",
			args: args{
				root1: &structs.TreeNode{
					Val: 1,
					Left: &structs.TreeNode{
						Val:   2,
						Left:  &structs.TreeNode{Val: 3, Left: nil, Right: nil},
						Right: nil,
					},
					Right: nil,
				},
				root2: &structs.TreeNode{
					Val:   1,
					Left:  &structs.TreeNode{Val: 3, Left: nil, Right: nil},
					Right: nil,
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leafSimilar(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("leafSimilar() = %v, want = %v", got, tt.want)
			}
		})
	}
}
