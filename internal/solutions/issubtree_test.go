package solutions

import (
	"testing"

	"github.com/therenotomorrow/leetcode/internal/structs"
)

func TestIsSubtree(t *testing.T) {
	type args struct {
		root    *structs.TreeNode
		subRoot *structs.TreeNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "smoke 1",
			args: args{
				root: &structs.TreeNode{
					Val: 3,
					Left: &structs.TreeNode{
						Val:   4,
						Left:  &structs.TreeNode{Val: 1, Left: nil, Right: nil},
						Right: &structs.TreeNode{Val: 2, Left: nil, Right: nil},
					},
					Right: &structs.TreeNode{Val: 5, Left: nil, Right: nil},
				},
				subRoot: &structs.TreeNode{
					Val:   4,
					Left:  &structs.TreeNode{Val: 1, Left: nil, Right: nil},
					Right: &structs.TreeNode{Val: 2, Left: nil, Right: nil},
				},
			},
			want: true,
		},
		{
			name: "smoke 2",
			args: args{
				root: &structs.TreeNode{
					Val: 3,
					Left: &structs.TreeNode{
						Val:  4,
						Left: &structs.TreeNode{Val: 1, Left: nil, Right: nil},
						Right: &structs.TreeNode{
							Val:   2,
							Left:  &structs.TreeNode{Val: 0, Left: nil, Right: nil},
							Right: nil,
						},
					},
					Right: &structs.TreeNode{Val: 5, Left: nil, Right: nil},
				},
				subRoot: nil,
			},
			want: false,
		},
		{
			name: "test 162: wrong answer",
			args: args{
				root: &structs.TreeNode{
					Val:   1,
					Left:  &structs.TreeNode{Val: 1, Left: nil, Right: nil},
					Right: nil,
				},
				subRoot: &structs.TreeNode{Val: 1, Left: nil, Right: nil},
			},
			want: true,
		},
		{
			name: "test 177: wrong answer",
			args: args{
				root: &structs.TreeNode{
					Val: 3,
					Left: &structs.TreeNode{
						Val:   4,
						Left:  &structs.TreeNode{Val: 1, Left: nil, Right: nil},
						Right: nil,
					},
					Right: &structs.TreeNode{
						Val:   5,
						Left:  &structs.TreeNode{Val: 2, Left: nil, Right: nil},
						Right: nil,
					},
				},
				subRoot: &structs.TreeNode{
					Val:   3,
					Left:  &structs.TreeNode{Val: 1, Left: nil, Right: nil},
					Right: &structs.TreeNode{Val: 2, Left: nil, Right: nil},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.args.root, tt.args.subRoot); got != tt.want {
				t.Errorf("isSubtree() = %v, want = %v", got, tt.want)
			}
		})
	}
}
