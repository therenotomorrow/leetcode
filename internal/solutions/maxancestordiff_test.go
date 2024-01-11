package solutions

import (
	"testing"

	"github.com/therenotomorrow/leetcode/internal/structs"
)

func TestMaxAncestorDiff(t *testing.T) {
	type args struct {
		root *structs.TreeNode
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1",
			args: args{
				root: &structs.TreeNode{
					Val: 8,
					Left: &structs.TreeNode{
						Val:  3,
						Left: &structs.TreeNode{Val: 1, Left: nil, Right: nil},
						Right: &structs.TreeNode{
							Val:   6,
							Left:  &structs.TreeNode{Val: 4, Left: nil, Right: nil},
							Right: &structs.TreeNode{Val: 7, Left: nil, Right: nil},
						},
					},
					Right: &structs.TreeNode{
						Val:  10,
						Left: nil,
						Right: &structs.TreeNode{
							Val:   14,
							Left:  &structs.TreeNode{Val: 13, Left: nil, Right: nil},
							Right: nil,
						},
					},
				},
			},
			want: 7,
		},
		{
			name: "smoke 2",
			args: args{
				root: &structs.TreeNode{
					Val:  1,
					Left: nil,
					Right: &structs.TreeNode{
						Val:  2,
						Left: nil,
						Right: &structs.TreeNode{
							Val:   0,
							Left:  &structs.TreeNode{Val: 3, Left: nil, Right: nil},
							Right: nil,
						},
					},
				},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAncestorDiff(tt.args.root); got != tt.want {
				t.Errorf("maxAncestorDiff() = %v, want = %v", got, tt.want)
			}
		})
	}
}
